package manager

import (
	"fmt"
	"strings"

	. "github.com/pspaces/gospace"
)

//CriarUser cria e insere um novo usuário em um ambiente
func CriarUser(espacoTupla *Space, ambName string, userName string) {
	espacoTupla.Put(ambName, userName)
	fmt.Printf("Um novo usuário foi criado no ambiente %s! Nome do usuário: %s \n", ambName, userName)
}

//ListarUser lista os usuários de um ambiente
func ListarUser(espacoTupla *Space, ambName string) {
	var userName string

	usuarios, _ := espacoTupla.QueryAll(ambName, &userName)

	fmt.Println("########## AMBIENTE ", ambName, " - USUÁRIOS ##########")
	for _, usuario := range usuarios {
		name := fmt.Sprintf("%v", usuario.GetFieldAt(1))
		if strings.Contains(name, "user") {
			fmt.Println(usuario.GetFieldAt(1))
		}
	}
}

//MoverUser muda o usuário de ambiente
func MoverUser(espacoTupla *Space, userName string, ambOrigem string, ambDestino string) {
	//Remove tupla ambiente de origem, usuário
	espacoTupla.GetP(ambOrigem, userName)

	//Cria tupla ambiente de destino, usuário
	espacoTupla.Put(ambDestino, userName)
	fmt.Printf("Usuário %s saiu do ambiente %s para o ambiente %s! \n", userName, ambOrigem, ambDestino)
}
