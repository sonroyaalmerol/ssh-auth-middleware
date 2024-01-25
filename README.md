# SSH Server Middleware 

This Go project serves as a proof-of-concept for an SSH server middleware with custom authentication (initially OIDC/OAuth2). The middleware generates a short-lived SSH certificate upon successful authentication and forwards the connection to an OpenSSH server. It demonstrates the integration of the golang.org/x/crypto/ssh package for SSH-related functionalities.

## Features

- **Authentication**: The middleware supports custom authentication such as OIDC/OAuth2.

- **SSH Certificate Generation**: Upon successful authentication, the middleware generates a short-lived SSH certificate. Customize the certificate generation logic in the `generateSSHCertificate` function in the `main.go` file.

- **Forwarding to OpenSSH Server**: After successful authentication and certificate generation, the middleware forwards the connection to a specified OpenSSH server. Configure the OpenSSH server address in the `forwardToOpenSSHServer` function in the `main.go` file.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**Note**: This project is a proof-of-concept, and for a production environment, additional security measures, error handling, and deployment considerations should be taken into account.
