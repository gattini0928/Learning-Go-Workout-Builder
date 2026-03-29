# REST API WORKOUT BUILDER - ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white) ![HTMX](https://img.shields.io/badge/HTMX-2E2E2E?style=for-the-badge&logo=htmx&logoColor=white)

Projeto utilizado para aprender a implementar API'S REST com Golang, pacote net/http nativo.

📌 Fluxo da Aplicação

🔐 Autenticação
  Criação de Conta → /signup → validação de dados → redirecionamento para homepage
  Login → /login → validação de credenciais → acesso à homepage
  (Futuro) Implementação de autenticação com JWT
  
🛠️ Administração
  Admin pode criar exercícios e disponibilizá-los em /exercises
  Admin pode deletar um exercício, apenas se ele não estiver vinculado a nenhum treino de usuário
  
🏋️ Funcionalidades do Usuário
  Usuário pode selecionar exercícios e montar um treino em /workouts
  Usuário pode:
    visualizar todos os seus treinos
    visualizar um treino específico
  Usuário pode editar:
    número de repetições
    número de séries de cada exercício dentro do treino
    
🚀 Funcionalidades Futuras
  Página de Packages (treinos completos):
    visualizar treinos prontos
    adicionar treinos completos aos seus próprios treinos
  Integração de frontend com HTMX

  ---

## 🚀 Como Rodar o Projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/gattini0928/Learning-Go-Workout-Builder.git
cd Learning-Go-Workout-Builder
```

### 2. Instalar depedências
```bash
go mod tidy
```

### 3. Configurar o arquivo .env
```bash
cp .env.example .env
```

### 4. Criar o banco de dados
```bash
CREATE DATABASE workout_builder;
```


### 5. Criar as tabelas do banco

Depois de criar o banco `workout_builder`, execute o arquivo `migrations/schema.sql`.

**Usando o terminal:**

```bash
psql -U postgres -d workout_builder -f schema.sql
```

### 6. Rodar a aplicação
Como o main.go está dentro da pasta cmd/api/, use o comando abaixo:
```bash
go run cmd/api/main.go
```

---

## Pontos de Melhoria Futura / Sugestões de Melhoria

Projeto encerrado intencionalmente para avançar no aprendizado — aberto a contribuições e melhorias.

Falta implementar os `pacotes(packages)`

**futuramente**: HTMX + Autenticação JWT 🚀
---

💻✌️ - `Dicas e críticas são sempre bem-vindas - Obrigado` ✌️💻