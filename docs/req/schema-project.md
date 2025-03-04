```
/blue-whale-platform                            # Contiene la plataforma Blue Whale
│── backend/                                    # Contiene el backend de la plataforma
│   ├── service/                                # Contiene los microservicios que forman el backend
│   │   ├── api-gateway/                        # Contiene el servicio api-gateway
│   │   ├── idp/                                # Contiene el servicio IdP
│   │   │   ├── cmd/                            # Contiene los puntos de entrada para la aplicación
│   │   │   │   ├── main.go                     # Punto de entrada principal del servicio
│   │   │   │   ├── server.go                   # Configuración y arranque del servidor HTTP
│   │   │   ├── internal/                       # Código específico del dominio del servicio, no accesible desde otros módulos
│   │   │   │   ├── domain/                     # Define los conceptos principales del negocio
│   │   │   │   │   ├── models/                 # Definición de entidades (User, Session, Token, etc.)
│   │   │   │   │   ├── repositories/           # Interfaces para persistencia de datos
│   │   │   │   │   ├── services/               # Lógica de negocio relacionada con autenticación y federación de identidad
│   │   │   │   ├── application/                # Casos de uso y lógica de aplicación
│   │   │   │   │   ├── usecases/               # Implementación de casos de uso
│   │   │   │   │   ├── ports/                  # Interfaces de entrada y salida para la comunicación con otros servicios
│   │   │   │   ├── infrastructure/             # Integraciones con tecnologías externas
│   │   │   │   │   ├── http/                   # Handlers y middlewares HTTP
│   │   │   │   │   ├── db/                     # Implementación de persistencia en PostgreSQL
│   │   │   │   │   ├── oauth/                  # Integración con OpenID Connect
│   │   │   │   │   ├── cache/                  # Implementación con Redis para tokens
│   │   │   │   │   ├── security/               # Cifrado y validación de credenciales
│   │   │   ├── pkg/                            # Paquetes reutilizables dentro del servicio
│   │   │   │   ├── logger/                     # Gestión centralizada de logs
│   │   │   │   ├── config/                     # Carga y validación de variables de entorno
│   │   │   │   ├── errors/                     # Manejo centralizado de errores
│   │   │   ├── deployments/                    # Contiene los archivos de configuración para el despliegue
│   │   │   │   ├── Dockerfile                  # Configuración del contenedor para IdP
│   │   │   │   ├── k8s/                        # YAMLs para despliegue en Kubernetes
│   │   │   │   ├── configmaps/                 # Variables de entorno en Kubernetes
│   │   │   ├── go.mod                          # Manejo de dependencias del servicio en Golang
│   │   │   ├── go.sum                          # Manejo de dependencias del servicio en Golang
│   │   │   ├── README.md                       # Documentación sobre la configuración y ejecución del servicio
│   │   ├── tenant-service/                     # Contiene el servicio tenat
│   │   ├── notification-service/               # Contiene el servicio notification
│   │   ├── monitoring/                         # Contiene el servicio monitoring
│   ├── messaging/                              #
│   │   ├── rabbitmq/                           # Configuración y bindings para la gestión de eventos en RabbitMQ
│   ├── storage/                                #
│   │   ├── database/                           # Base de datos PostgreSQL. Scripts de migración y gestión de esquemas
│   ├── test/                                   #
│   │   ├── unit/                               # Pruebas unitarias para cada microservicio
│   │   │   ├── api-gateway/                    #
│   │   │   ├── idp/                            #
│   │   │   ├── tenat-service/                  #
│   │   │   ├── notification-service/           #
│   │   ├── integration/                        # Pruebas de integración entre componentes internos y bases de datos
│   │   │   ├── api-gateway/                    #
│   │   │   ├── idp/                            #
│   │   │   ├── tenat-service/                  #
│   │   │   ├── notification-service/           #
│   │   ├── features/                           # Pruebas de aceptación y comportamiento
│   │   │   ├── api-gateway/                    #
│   │   │   ├── idp/                            #
│   │   │   ├── tenat-service/                  #
│   │   │   ├── notification-service/           #
│   ├── go.work
│   ├── go.work.sum
│── frontend/
│   ├── web/
│   ├── mobile/
│   │   ├── ios/
│   │   ├── android/
│── devops/
│   ├── ci-cd/
│   │   ├── github-actions/
│   ├── docker/
│   │   ├── Dockerfile
│   │   ├── docker-compose.yaml
│   ├── k8s/
│   │   ├── deployments/
│   │   ├── services/
│   │   ├── config-maps/
│   ├── terraform/
│── docs/
│   ├── api-specs/
│   ├── terraform/

```