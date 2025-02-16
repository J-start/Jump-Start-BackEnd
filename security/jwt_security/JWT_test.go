package jwt_security

import (
	"testing"
)

func TestValidToken(t *testing.T) {
	message := "testEmail@gmail.com"
	token, err := GenerateToken(message)
	if err != nil {
		t.Fatalf("Erro ao gerar token: %v", err)
	}

	decodedMessage, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Erro ao validar token: %v", err)
	}

	if decodedMessage.UserEmail != message {
		t.Errorf("Mensagem decodificada não corresponde. Esperado: %s, Recebido: %s", message, decodedMessage)
	}
}

func TestInvalidToken(t *testing.T) {
	invalidToken := "token iválido"
	_, err := ValidateToken(invalidToken)
	if err == nil {
		t.Error("Token inválido foi considerado válido")
	}
}

func TestValidTokenWith5Minutes(t *testing.T) {
	message := "myemail@gmail.com"
	token, err := GenerateTokenWithNMinutes(message,5)
	if err != nil {
		t.Error("Token válido foi considerado inválido")
	}
	decodedMessage, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Erro ao validar token: %v", err)
	}
	if decodedMessage.UserEmail != message {
		t.Errorf("Mensagem decodificada não corresponde. Esperado: %s, Recebido: %s", message, decodedMessage)
	}
}

func TestInValidTokenWithNMinutes(t *testing.T) {
	message := "myemail@gmail.com"
	token, err := GenerateTokenWithNMinutes(message,-1)
	if err == nil {
		t.Error("Token válido foi considerado inválido")
	}
	_, err2 := ValidateToken(token)
	if err2 == nil {
		t.Fatalf("Token deveria ser inválido: %v", err2)
	}

}
func TestTamperedToken(t *testing.T) {
	message := "testEmail@gmail.com"
	token, err := GenerateToken(message)
	if err != nil {
		t.Fatalf("Erro ao gerar token: %v", err)
	}

	tamperedToken := token + "modificado"
	_, err = ValidateToken(tamperedToken)
	if err == nil {
		t.Error("Token com assinatura alterada foi considerado válido")
	}
}



