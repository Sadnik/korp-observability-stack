# Demonstração de Stack de Observabilidade

Plataforma de observabilidade construída com Go, Nginx, Prometheus, Grafana, Docker Compose e Ansible.

O projeto demonstra implantação de aplicações containerizadas, automação de infraestrutura, coleta de métricas, provisionamento de dashboards e validação de desempenho sob carga.

Testado com **101.219 requisições por segundo**, **0 falhas** e **latência P99 de 15,05ms** utilizando **400 conexões simultâneas**.

## Capturas de Tela

### Dashboard do Grafana
![Dashboard do Grafana](https://github.com/user-attachments/assets/6e019ed7-1316-44bf-b1b9-352b20a492b9)

### Targets do Prometheus
![Targets do Prometheus](https://github.com/user-attachments/assets/990c2b4f-0101-459b-a047-a172a24e043d)

### Taxa de Requisições no Prometheus
![Taxa de Requisições](https://github.com/user-attachments/assets/6cd38172-94d9-40e7-b1d8-f95de38a9a54)

---

## Arquitetura

```text
                           +----------------+
                           |     Client     |
                           +--------+-------+
                                    |
                                    v
                           +----------------+
                           |     Nginx      |
                           |      :80       |
                           +--------+-------+
                                    |
                                    v
                           +----------------+
                           |   Go Service   |
                           |     :8080      |
                           +--------+-------+
                                    |
                     +---------------+---------------+
                     |                               |
                     v                               v
            +----------------+              +----------------+
            |  Prometheus    | -----------> |    Grafana     |
            |     :9090      |              |     :3000      |
            +----------------+              +----------------+
```

Todos os serviços executam em uma rede Docker Bridge isolada.

O Nginx atua como ponto único de entrada da aplicação. O serviço Go expõe métricas para o Prometheus, que realiza a coleta periódica dos dados e os disponibiliza para visualização no Grafana.

---

## Stack Tecnológica

| Componente     | Função                               |
| -------------- | ------------------------------------ |
| Go 1.21        | API HTTP e instrumentação Prometheus |
| Nginx          | Proxy reverso e roteamento           |
| Prometheus     | Coleta e armazenamento de métricas   |
| Grafana        | Visualização e dashboards            |
| Docker Compose | Orquestração dos containers          |
| Ansible        | Automação de implantação             |

---

## Inicialização Rápida

### Subir a stack completa

```bash
docker compose up -d --build
```

### Validar funcionamento

```bash
curl http://localhost/demo-service
```

---

## Endpoints

| Endpoint          | Descrição                                     |
| ----------------- | --------------------------------------------- |
| GET /demo-service | Retorna resposta JSON contendo nome e horário |
| GET /health       | Endpoint de verificação de saúde do Nginx     |
| GET /metrics      | Métricas utilizadas pelo Prometheus           |
| :9090             | Interface Web do Prometheus                   |
| :3000             | Interface Web do Grafana                      |

Credenciais padrão do Grafana:

```text
admin / admin
```

---

## Implantação com Ansible

O playbook automatiza:

* Instalação do Docker
* Validação do ambiente
* Implantação da aplicação
* Inicialização dos serviços
* Verificação de saúde da stack

Execução:

```bash
ansible-playbook ansible/playbook.yml
```

Esse comando funciona a partir da raiz do repositório porque `ansible.cfg` define o inventário, conexão local e elevação de privilégios por padrão.

---

## Teste de Desempenho

Benchmark executado com:

```bash
wrk -t16 -c400 -d30s --latency http://localhost/demo-service
```

### Resultados

| Métrica                 | Valor       |
| ----------------------- | ----------- |
| Requisições por segundo | 101.219     |
| Total de requisições    | 3.046.709   |
| Falhas                  | 0           |
| Latência P50            | 3,62ms      |
| Latência P75            | 5,81ms      |
| Latência P90            | 8,54ms      |
| Latência P99            | 15,05ms     |
| Latência máxima         | 70,99ms     |
| Conexões simultâneas    | 400         |
| Duração                 | 30 segundos |

---

## Métricas Coletadas

| Métrica                       | Tipo      | Descrição                                          |
| ----------------------------- | --------- | -------------------------------------------------- |
| http_requests_total           | Counter   | Total de requisições por método, endpoint e status |
| http_request_duration_seconds | Histogram | Distribuição de latência das requisições           |
| service_availability          | Gauge     | Disponibilidade do serviço                         |

Os dashboards do Grafana apresentam:

* Taxa de requisições (5m)
* Latência P95
* Latência P99
* Disponibilidade
* Total de requisições 

---

## Decisões de Arquitetura

### Camada de Proxy Reverso

O Nginx é o único componente exposto externamente. O serviço Go permanece acessível apenas pela rede interna do Docker.

### Isolamento de Métricas

O endpoint `/metrics` é utilizado exclusivamente pelo Prometheus e não fica disponível externamente através do proxy.

### Health Check Independente

O endpoint `/health` é servido diretamente pelo Nginx, permitindo validar a camada de proxy independentemente da aplicação.

### Infraestrutura como Código

Toda a configuração da stack é versionada e reproduzível através de Docker Compose e Ansible.

### Provisionamento Automático do Grafana

Datasources e dashboards são carregados automaticamente durante a inicialização.

### Build Multi-Stage

A imagem final contém apenas o binário compilado e os componentes necessários para execução.

### Timeouts Explícitos

ReadTimeout, WriteTimeout e IdleTimeout foram configurados para evitar conexões presas e aumentar a resiliência da aplicação.

### Observabilidade Desde o Início

Métricas de desempenho, disponibilidade e utilização da aplicação ficam disponíveis imediatamente após a implantação.