package manager

import (
	"fmt"

	. "github.com/pspaces/gospace"
)

//CriarAmb inicia um novo ambiente no Espaco de Tuplas
func CriarAmb(espacoTupla *Space, ambName string) {
	espacoTupla.Put(ambName)
	fmt.Println("Um novo ambiente foi criado! Nome do ambiente: ", ambName)
}

//ListarAmb lista os ambientes existentes
func ListarAmb(espacoTupla *Space) {
	var ambName string
	var name string

	tupla, _ := espacoTupla.QueryAll(&ambName)

	fmt.Println("########## AMBIENTES ##########")
	for _, elemento := range tupla {
		integrantes, _ := espacoTupla.QueryAll(elemento.GetFieldAt(0), &name)

		fmt.Printf("\n [%s] \n", elemento.GetFieldAt(0))
		fmt.Printf("- Número de conectados: %d  \n", len(integrantes))
		for _, integrante := range integrantes {
			fmt.Println("**** ", integrante.GetFieldAt(1))
		}
	}
	fmt.Println()
}

//RemoverAmbs remove ambientes vazios
func RemoverAmbs(espacoTupla *Space) {
	var ambName string
	var strName string
	ambientes, _ := espacoTupla.QueryAll(&ambName)
	for _, ambiente := range ambientes {
		//Verifica se existem dispositivos ou usuários no ambiente
		_, err := espacoTupla.QueryP(ambiente.GetFieldAt(0), &strName)

		if err != nil {
			name := ambiente.GetFieldAt(0)
			espacoTupla.GetP(name)
			fmt.Println("O seguinte ambiente foi destruído: ", name)
		}
	}
}
