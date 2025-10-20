---
title: "Validação e Formatação de Dados Críticos - MVP v1.0"
version: "1.0"
date_created: "2025-10-16"
last_updated: "2025-10-20"
author: "Desenvolvedor"
status: "Planned"
priority: "Alta"
tags: ["feature", "validation", "data-integrity", "critical"]
project: "MindTrace"
---

# ✅ Validação e Formatação de Dados Críticos - MVP v1.0

![Status: Planned](https://img.shields.io/badge/status-Planned-blue) ![Priority: High](https://img.shields.io/badge/priority-High-red) ![Timeline: 7-11 days](https://img.shields.io/badge/timeline-7--11%20days-green)

## 📊 Visão Geral

Plano simplificado para implementar validação robusta de dados críticos da aplicação MindTrace. Foca em validações de entrada (email, CPF, datas, atividades de autocuidado) com verificações tanto no frontend quanto no backend.

**📊 Escopo**: 1 desenvolvedor | Projeto acadêmico | Iterativo

## 1️⃣ Requisitos Críticos

- **REQ-001**: Email - validação de formato + unicidade no BD
- **REQ-002**: CPF - validação de formato, dígitos e unicidade
- **REQ-003**: Data de nascimento - não futuro, maior de 8 anos
- **REQ-004**: Telefone - formato brasileiro (opcional mas validado se preenchido)
- **REQ-005**: Atividades de autocuidado - salvar como array JSON estruturado
- **REQ-006**: Validação do formulário de Registro de Humor - ranges corretos, data válida

## 2️⃣ Fases de Implementação

### Fase 1: Backend - Validadores (1-2 dias) 🔍

**Objetivo**: Criar pacote centralizado de funções de validação

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-001 | Criar `backend/interno/pkg/validadores/email.go` - ValidarEmail(string) bool | ⏳ |
| TASK-002 | Criar `backend/interno/pkg/validadores/cpf.go` - ValidarCPF, NormalizarCPF | ⏳ |
| TASK-003 | Criar `backend/interno/pkg/validadores/data.go` - ValidarDataNascimento, ValidarDataNaoFutura | ⏳ |
| TASK-004 | Criar `backend/interno/pkg/validadores/telefone.go` - ValidarTelefone, NormalizarTelefone | ⏳ |
| TASK-005 | Adicionar validações em DTOs com tags `binding:"required,email"` etc | ⏳ |

### Fase 2: Backend - Controladores (2-3 dias) 🔐

**Objetivo**: Aplicar validação ao fluxo de criação/edição

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-006 | Atualizar endpoint registro de usuário (email, CPF, data) | ⏳ |
| TASK-007 | Atualizar endpoint edição de perfil | ⏳ |
| TASK-008 | Adicionar validações ao registro de humor (ranges, data) | ⏳ |
| TASK-009 | Criar middleware de erro padronizado para validação | ⏳ |

### Fase 3: Frontend - Validações (2-3 dias) 🎨

**Objetivo**: Feedback visual ao usuário antes de enviar

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-010 | Adicionar validações no formulário de cadastro (RegistroHumor.vue) | ⏳ |
| TASK-011 | Adicionar validações no formulário de edição de perfil | ⏳ |
| TASK-012 | Mostrar mensagens de erro com toast/alert | ⏳ |
| TASK-013 | Desabilitar submit enquanto há erros | ⏳ |

### Fase 4: Armazenamento de Autocuidado (1-2 dias) 💾

**Objetivo**: Converter autocuidado de string para JSON estruturado

| Task | Descrição | Status |
|------|-----------|--------|
| TASK-014 | Criar migration SQL: alterar coluna `auto_cuidado` para JSONB/JSON | ⏳ |
| TASK-015 | Atualizar struct `RegistroHumor` com tipo correto | ⏳ |
| TASK-016 | Atualizar serialização/desserialização no controlador | ⏳ |
| TASK-017 | Atualizar frontend para enviar array estruturado | ⏳ |

## 3️⃣ Arquivos a Modificar/Criar

### Backend

**Novos arquivos**
- ✨ `backend/interno/pkg/validadores/` - Novo pacote com 5 arquivos
  - `email.go`
  - `cpf.go`
  - `data.go`
  - `telefone.go`
  - `autocuidado.go`

**Arquivos modificados**
- 🔄 `backend/interno/aplicacao/controladores/` - Atualizar controladores
- 🔄 `backend/interno/dominio/registro_humor.go` - Atualizar tipo AutoCuidado
- 🔄 `backend/cmd/api/main.go` - Se needed, adicionar imports

### Frontend

**Arquivos modificados**
- 🔄 `frontend/src/views/dashboard-paciente/RegistroHumor.vue` - Adicionar validações
- 🔄 `frontend/src/services/api.js` - Sem mudanças (apenas trata erros)
- 🔄 Formulários de cadastro/edição - Adicionar validações

### Database

**Migrations**
- 🔄 `schema.sql` / Migration SQL - Atualizar tipo de auto_cuidado

## 4️⃣ Exemplos de Código

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

## 5️⃣ Testes Básicos

| # | Teste | Esperado |
|---|-------|----------|
| TEST-001 | ValidarCPF('123.456.789-09') | ✓ retorna false (CPF inválido) |
| TEST-002 | ValidarEmail('user@domain.com') | ✓ retorna true |
| TEST-003 | ValidarDataNascimento(data_futura) | ✓ retorna false |
| TEST-004 | POST /pacientes/registrar com email inválido | ✓ retorna 400 |
| TEST-005 | POST /registro-humor com humor fora do range | ✓ retorna 400 |

## 6️⃣ Riscos & Mitigações

| Risk | Impacto | Mitigação |
|------|--------|-----------|
| Migration de dados pode impactar registros antigos | Alto | Backup recomendado |
| Validações muito rígidas podem rejeitar dados válidos | Médio | Testar bem antes de deploy |
| Dados antigos inconsistentes podem causar falha na migration | Alto | Validação prévia de dados existentes |

## 7️⃣ Timeline Estimada

| Fase | Dias | Acumulado |
|------|------|-----------|
| 1: Backend - Validadores | 1-2 | 1-2 |
| 2: Backend - Controladores | 2-3 | 3-5 |
| 3: Frontend - Validações | 2-3 | 5-8 |
| 4: Armazenamento Autocuidado | 1-2 | 6-10 |
| **TOTAL** | **7-11 dias** | **7-11 dias** |

## 8️⃣ Ordem Recomendada

1. **Primeiro**: Criar validadores (Fase 1) - independentes
2. **Depois**: Aplicar nos controladores (Fase 2) - backend precisa estar pronto
3. **Depois**: Adicionar feedback no frontend (Fase 3) - após API estar validando
4. **Último**: Migração de autocuidado (Fase 4) - pode ser feita depois se necessário

---

## 📌 Metadados

| Campo | Valor |
|-------|-------|
| **Criado em** | 2025-10-16 |
| **Atualizado em** | 2025-10-20 |
| **Duração estimada** | 7-11 dias |
| **Prioridade** | Alta ⬆️ |
| **Contexto** | Disciplina Engenharia de Software - Projeto MindTrace |
| **Status** | Planned 🔵 |
