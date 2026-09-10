package password

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {

	// Gera um salt aleatório de 16 bytes
	salt := make([]byte, 16)
	_, err := rand.Read(salt)

	// Verifica se ocorreu algum erro ao gerar o salt
	if err != nil {
		return "", err
	}

	// Configurações do Argon2
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLength := uint32(32)

	// Hash o password com o salt usando Argon2id
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		time,
		memory,
		threads,
		keyLength,
	)

	// Codifica o salt e o hash em base64 para armazenar
	// como string evitando erro de caracteres quebrados e bytes incorretos
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)

	// Retorna o salt e o hash
	return saltEncoded + ":" + hashEncoded, nil
}
