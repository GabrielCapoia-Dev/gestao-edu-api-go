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

	// Hash o password com o salt
	/*
		[]byte(password) → senha em bytes
		salt             → salt aleatório
		1                → número de iterações
		64*1024          → memória: 64 MiB
		4                → paralelismo
		32               → tamanho final do hash: 32 bytes
	*/
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Codifica o salt e o hash em base64 para armazenar como string evitando erro de caracteres quebrados e bytes incorretos
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)

	// Retorna o salt e o hash
	return saltEncoded + ":" + hashEncoded, nil
}

