package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// IPRateLimiter struct para gerenciar limites por IP
type IPRateLimiter struct {
	ips    map[string]*client
	mu     sync.Mutex
	rate   float64 // tokens por segundo
	burst  int     // tamanho maximo do balde
}

type client struct {
	tokens     float64
	lastUpdate time.Time
}

// NewIPRateLimiter cria um novo limitador
func NewIPRateLimiter(rate float64, burst int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips:   make(map[string]*client),
		rate:  rate,
		burst: burst,
	}

	// Limpeza automatica de clientes antigos a cada minuto
	go func() {
		for {
			time.Sleep(time.Minute)
			i.mu.Lock()
			for ip, c := range i.ips {
				if time.Since(c.lastUpdate) > 3*time.Minute {
					delete(i.ips, ip)
				}
			}
			i.mu.Unlock()
		}
	}()

	return i
}

// Allow verifica se o IP pode fazer a requisicao
func (i *IPRateLimiter) Allow(ip string) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = &client{
			tokens:     float64(i.burst),
			lastUpdate: time.Now(),
		}
		i.ips[ip] = limiter
	}

	now := time.Now()
	elapsed := now.Sub(limiter.lastUpdate).Seconds()
	limiter.lastUpdate = now

	// Adiciona tokens baseados no tempo passado
	limiter.tokens += elapsed * i.rate
	if limiter.tokens > float64(i.burst) {
		limiter.tokens = float64(i.burst)
	}

	if limiter.tokens >= 1 {
		limiter.tokens--
		return true
	}

	return false
}

// RateLimitMiddleware cria o middleware para o Gin
func RateLimitMiddleware(rate float64, burst int) gin.HandlerFunc {
	limiter := NewIPRateLimiter(rate, burst)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !limiter.Allow(ip) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"erro": "Muitas requisições. Tente novamente mais tarde.",
			})
			return
		}
		c.Next()
	}
}
