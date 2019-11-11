package manager

import (
	"fmt"
	"strings"

	. "github.com/pspaces/gospace"
)

//CriarDisp cria e insere um novo dispositivo em um ambiente
func CriarDisp(espacoTupla *Space, ambName string, dispName string) {
	espacoTupla.Put(ambName, dispName)
	fmt.Printf("Um novo dispositivo foi criado no ambiente %s! Nome do dispositivo: %s \n", ambName, dispName)
}

//ListarDisp lista os dispositivos de um ambiente
func ListarDisp(espacoTupla *Space, ambName string) {
	var dispName string

	dispositivos, _ := espacoTupla.QueryAll(ambName, &dispName)

	fmt.Println("########## AMBIENTE ", ambName, " - DISPOSITIVOS ##########")
	for _, dispositivo := range dispositivos {
		name := fmt.Sprintf("%v", dispositivo.GetFieldAt(1))
		if strings.Contains(name, "disp") {
			fmt.Println(dispositivo.GetFieldAt(1))
		}
	}
}

//MoverDisp muda o dispositivo de ambiente
func MoverDisp(espacoTupla *Space, dispName string, ambOrigem string, ambDestino string) {
	//Remove tupla ambiente de origem, dispositivo
	espacoTupla.GetP(ambOrigem, dispName)

	//Cria tupla ambiente de destino, dispositivo
	espacoTupla.Put(ambDestino, dispName)
	fmt.Printf("Dispositivo %s foi movido do ambiente %s para o ambiente %s! \n", dispName, ambOrigem, ambDestino)
}
