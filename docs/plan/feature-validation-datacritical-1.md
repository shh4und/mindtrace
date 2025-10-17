---
goal: "Validação e Formatação de Dados Críticos"
version: "1.0"
date_created: "2025-10-16"
owner: "Desenvolvedor"
status: 'Planned'
tags: ['validation', 'data-integrity', 'feature']
---

# Introdução

![Status: Planned](https://img.shields.io/badge/status-Planned-blue)

Plano simplificado para implementar validação robusta de dados críticos da aplicação MindTrace. Foca em validações de entrada (email, CPF, datas, atividades de autocuidado) com verificações tanto no frontend quanto no backend.

**Escopo**: 1 desenvolvedor | Projeto acadêmico | Iterativo

## 1. Requisitos Críticos

- **REQ-001**: Email - validação de formato + unicidade no BD
- **REQ-002**: CPF - validação de formato, dígitos e unicidade
- **REQ-003**: Data de nascimento - não futuro, maior de 8 anos
- **REQ-004**: Telefone - formato brasileiro (opcional mas validado se preenchido)
- **REQ-005**: Atividades de autocuidado - salvar como array JSON estruturado
- **REQ-006**: Validação do formulário de Registro de Humor - ranges corretos, data válida

## 2. Implementação

### Fase 1: Backend - Validadores (1-2 dias)

**Objetivo**: Criar pacote centralizado de funções de validação

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-001 | Criar `backend/interno/pkg/validadores/email.go` - ValidarEmail(string) bool | ⏳ |
| TASK-002 | Criar `backend/interno/pkg/validadores/cpf.go` - ValidarCPF, NormalizarCPF | ⏳ |
| TASK-003 | Criar `backend/interno/pkg/validadores/data.go` - ValidarDataNascimento, ValidarDataNaoFutura | ⏳ |
| TASK-004 | Criar `backend/interno/pkg/validadores/telefone.go` - ValidarTelefone, NormalizarTelefone | ⏳ |
| TASK-005 | Adicionar validações em DTOs com tags `binding:"required,email"` etc | ⏳ |

### Fase 2: Backend - Controladores (2-3 dias)

**Objetivo**: Aplicar validação ao fluxo de criação/edição

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-006 | Atualizar endpoint registro de usuário (email, CPF, data) | ⏳ |
| TASK-007 | Atualizar endpoint edição de perfil | ⏳ |
| TASK-008 | Adicionar validações ao registro de humor (ranges, data) | ⏳ |
| TASK-009 | Criar middleware de erro padronizado para validação | ⏳ |

### Fase 3: Frontend - Validações (2-3 dias)

**Objetivo**: Feedback visual ao usuário antes de enviar

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-010 | Adicionar validações no formulário de cadastro (RegistroHumor.vue) | ⏳ |
| TASK-011 | Adicionar validações no formulário de edição de perfil | ⏳ |
| TASK-012 | Mostrar mensagens de erro com toast/alert | ⏳ |
| TASK-013 | Desabilitar submit enquanto há erros | ⏳ |

### Fase 4: Armazenamento de Autocuidado (1-2 dias)

**Objetivo**: Converter autocuidado de string para JSON estruturado

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-014 | Criar migration SQL: alterar coluna `auto_cuidado` para JSONB/JSON | ⏳ |
| TASK-015 | Atualizar struct `RegistroHumor` com tipo correto | ⏳ |
| TASK-016 | Atualizar serialização/desserialização no controlador | ⏳ |
| TASK-017 | Atualizar frontend para enviar array estruturado | ⏳ |

## 3. Arquivos a Modificar/Criar

**Backend**
- ✨ `backend/interno/pkg/validadores/` - Novo pacote com 5 arquivos
- 🔄 `backend/interno/aplicacao/controladores/` - Atualizar controladores
- 🔄 `backend/interno/dominio/registro_humor.go` - Atualizar tipo AutoCuidado
- 🔄 `backend/cmd/api/main.go` - Se needed, adicionar imports

**Frontend**
- 🔄 `frontend/src/views/dashboard-paciente/RegistroHumor.vue` - Adicionar validações
- 🔄 `frontend/src/services/api.js` - Sem mudanças (apenas trata erros)
- 🔄 Formulários de cadastro/edição - Adicionar validações

**Database**
- 🔄 `schema.sql` / Migration SQL - Atualizar tipo de auto_cuidado

## 4. Código de Exemplo

### Backend - Validadores

```go
// backend/interno/pkg/validadores/email.go
package validadores

import "regexp"

func ValidarEmail(email string) bool {
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    re := regexp.MustCompile(pattern)
    return re.MatchString(email)
}

// backend/interno/pkg/validadores/cpf.go
package validadores

import (
    "regexp"
)

func ValidarCPF(cpf string) bool {
    cpf = NormalizarCPF(cpf)
    if len(cpf) != 11 {
        return false
    }
    // Verificar se todos os dígitos são iguais
    if cpf[0] == cpf[1] && cpf[1] == cpf[2] {
        return false
    }
    // Calcular dígitos verificadores
    return calcularDigito(cpf, 9) && calcularDigito(cpf, 10)
}

func NormalizarCPF(cpf string) string {
    pattern := regexp.MustCompile(`\D`)
    return pattern.ReplaceAllString(cpf, "")
}

func calcularDigito(cpf string, posicao int) bool {
    // Implementar algoritmo de cálculo
    return true
}

// backend/interno/pkg/validadores/data.go
package validadores

import "time"

func ValidarDataNascimento(data time.Time, idadeMinima int) bool {
    agora := time.Now()
    idade := agora.Year() - data.Year()
    if agora.Month() < data.Month() || (agora.Month() == data.Month() && agora.Day() < data.Day()) {
        idade--
    }
    return idade >= idadeMinima
}

func ValidarDataNaoFutura(data time.Time) bool {
    return data.Before(time.Now()) || data.Equal(time.Now())
}

// backend/interno/pkg/validadores/telefone.go
package validadores

import "regexp"

func ValidarTelefone(telefone string) bool {
    pattern := `^\(?[0-9]{2}\)?9?[0-9]{4}-?[0-9]{4}$`
    re := regexp.MustCompile(pattern)
    return re.MatchString(telefone)
}

func NormalizarTelefone(telefone string) string {
    pattern := regexp.MustCompile(`\D`)
    return pattern.ReplaceAllString(telefone, "")
}
```

### Frontend - Validação em Vue

```vue
<script setup>
import { ref, computed } from 'vue';

const email = ref('');
const emailError = computed(() => {
  if (!email.value) return '';
  const pattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  return pattern.test(email.value) ? '' : 'Email inválido';
});

const isFormValid = computed(() => !emailError.value && email.value);
</script>

<template>
  <input v-model="email" type="email" />
  <p v-if="emailError" class="text-red-600">{{ emailError }}</p>
  <button :disabled="!isFormValid">Enviar</button>
</template>
```

## 5. Testes Básicos

- **TEST-001**: ValidarCPF('123.456.789-09') deve retornar false (CPF inválido)
- **TEST-002**: ValidarEmail('user@domain.com') deve retornar true
- **TEST-003**: ValidarDataNascimento(data_futura) deve retornar false
- **TEST-004**: POST /pacientes/registrar com email inválido retorna 400
- **TEST-005**: POST /registro-humor com humor fora do range (1-5) retorna 400

## 6. Riscos

- **RISK-001**: Migration de dados pode impactar registros antigos (backup recomendado)
- **RISK-002**: Validações muito rígidas podem rejeitar dados válidos (testar bem)
- **RISK-003**: Se houver dados antigos inconsistentes, migration pode falhar

## 7. Estimativa

- **Total**: 7-11 dias de trabalho (com testes)
- **Prioridade**: Alta (dados críticos)

## 8. Ordem Recomendada

1. **Primeiro**: Criar validadores (Fase 1) - são independentes
2. **Depois**: Aplicar nos controladores (Fase 2) - backend precisa estar pronto
3. **Depois**: Adicionar feedback no frontend (Fase 3) - após API estar validando
4. **Último**: Migração de autocuidado (Fase 4) - pode ser feita depois se necessário

---

Criado: 2025-10-16
Simplificado para: 1 desenvolvedor | Projeto acadêmico
