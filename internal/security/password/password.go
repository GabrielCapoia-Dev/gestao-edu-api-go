package password

import (
	"golang.org/x/crypto/argon2"
	"crypto/rand"
)
func HashPassword(password string) (string, error) {

	// Gera um salt aleatório de 16 bytes
	salt := make([]byte, 16)
	_, err := rand.Read(salt)

	// Verifica se ocorreu algum erro ao gerar o salt
	if err != nil {
		return "", err
	}

	// Hash o password com o salt
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Retorna o salt e o hash
	return string(salt) + ":" + string(hash), nil
}

