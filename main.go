package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	. "github.com/pspaces/gospace"
	"github.com/ryllari/tuplespace/manager"
)

func main() {

	//Cria uma Tuple Space
	espacoTupla := NewSpace("Manager")
	ambcount := 1
	usercount := 1
	dispcount := 1

	var ambName, userName, dispName string

	for {
		var option string
		fmt.Println("#########################################################")
		fmt.Println("############   GERENCIAMENTO DE AMBIENTES   #############")
		fmt.Println("#########################################################")
		fmt.Println()
		fmt.Println("# DIGITE O NUMERO REFERENTE AO QUE DESEJA NO MENU")
		fmt.Println("1 - Criar ambiente")
		fmt.Println("2 - Listar ambientes")
		fmt.Println("3 - Cadastrar usuário")
		fmt.Println("4 - Listar usuários de um ambiente")
		fmt.Println("5 - Mover usuários de um ambiente")
		fmt.Println("6 - Cadastrar dispositivo")
		fmt.Println("7 - Listar dispositivos de um ambiente")
		fmt.Println("8 - Mover dispositivos de um ambiente")
		fmt.Println("9 - Destruir ambientes vazios")
		fmt.Println("0 - Sair")
		fmt.Print("Digite uma opção: ")
		fmt.Scanf("%s", &option)
		fmt.Printf("\n\n")

		switch option {
		case "1":
			fmt.Println("Criando novo ambiente...")
			ambName = "amb" + strconv.Itoa(ambcount)
			manager.CriarAmb(&espacoTupla, ambName)
			ambcount++
		case "2":
			fmt.Println("Listando ambientes...")
			manager.ListarAmb(&espacoTupla)
		case "3":
			fmt.Printf("\nEm qual ambiente deseja cadastrar o usuário? ")
			fmt.Scanf("%s", &ambName)
			_, err := espacoTupla.QueryP(ambName)
			if err == nil {
				fmt.Println("Cadastrando um novo usuário...")
				userName := "user" + strconv.Itoa(usercount)
				manager.CriarUser(&espacoTupla, ambName, userName)
				usercount++
			} else {
				fmt.Printf("\nNão foi possível cadastrar usuário: O ambiente %s não existe!\n", ambName)
			}
		case "4":
			fmt.Printf("\nDe qual ambiente deseja listar usuários? ")
			fmt.Scanf("%s", &ambName)
			_, err := espacoTupla.QueryP(ambName)
			if err == nil {
				manager.ListarUser(&espacoTupla, ambName)
			} else {
				fmt.Printf("\nNão foi possível listar usuários: O ambiente %s não existe!\n", ambName)
			}
		case "5":
			var ambOrigem, ambDestino string
			fmt.Printf("\nDe qual ambiente de origem do usuário desejado? ")
			fmt.Scanf("%s", &ambOrigem)
			_, err := espacoTupla.QueryP(ambOrigem)
			if err == nil {
				fmt.Printf("\nQual o usuário desejado? ")
				fmt.Scanf("%s", &userName)
				_, err = espacoTupla.QueryP(ambOrigem, userName)
				if err != nil {
					fmt.Printf("\nDe qual ambiente de destino do usuário desejado? ")
					fmt.Scanf("%s", &ambDestino)
					_, err := espacoTupla.QueryP(ambDestino)
					if err == nil {
						fmt.Println("Movendo usuário...")
						manager.MoverUser(&espacoTupla, userName, ambOrigem, ambDestino)
					} else {
						fmt.Printf("\nNão foi possível mover usuário: O ambiente %s não existe!\n", ambDestino)
					}
				} else {
					fmt.Printf("\nNão foi possível mover usuário: O usuário %s não existe no ambiente %s!\n", userName, ambOrigem)
				}
				manager.ListarUser(&espacoTupla, ambName)
			} else {
				fmt.Printf("\nNão foi possível mover usuário: O ambiente %s não existe!\n", ambOrigem)
			}
		case "6":
			fmt.Printf("\nEm qual ambiente deseja cadastrar o dispositivo? ")
			fmt.Scanf("%s", &ambName)
			_, err := espacoTupla.QueryP(ambName)
			if err == nil {
				fmt.Println("Cadastrando um novo dispositivo...")
				dispName := "disp" + strconv.Itoa(dispcount)
				manager.CriarDisp(&espacoTupla, ambName, dispName)
				dispcount++
			} else {
				fmt.Printf("\nNão foi possível cadastrar dispositivo: O ambiente %s não existe!\n", ambName)
			}
		case "7":
			fmt.Printf("\nDe qual ambiente deseja listar dispositivos? ")
			fmt.Scanf("%s", &ambName)
			_, err := espacoTupla.QueryP(ambName)
			if err == nil {
				manager.ListarDisp(&espacoTupla, ambName)
			} else {
				fmt.Printf("\nNão foi possível listar dispositivos: O ambiente %s não existe!\n", ambName)
			}
		case "8":
			var ambOrigem, ambDestino string
			fmt.Printf("\nDe qual ambiente de origem do dispositivo desejado? ")
			fmt.Scanf("%s", &ambOrigem)
			_, err := espacoTupla.QueryP(ambOrigem)
			if err == nil {
				fmt.Printf("\nQual o dispositivo desejado? ")
				fmt.Scanf("%s", &dispName)
				_, err = espacoTupla.QueryP(ambOrigem, dispName)
				if err != nil {
					fmt.Printf("\nDe qual ambiente de destino do dispositivo desejado? ")
					fmt.Scanf("%s", &ambDestino)
					_, err := espacoTupla.QueryP(ambDestino)
					if err == nil {
						fmt.Println("Movendo dispositivo...")
						manager.MoverUser(&espacoTupla, dispName, ambOrigem, ambDestino)
					} else {
						fmt.Printf("\nNão foi possível mover dispositivo: O ambiente %s não existe!\n", ambDestino)
					}
				} else {
					fmt.Printf("\nNão foi possível mover dispositivo: O dispositivo %s não existe no ambiente %s!\n", dispName, ambOrigem)
				}
				manager.ListarUser(&espacoTupla, ambName)
			} else {
				fmt.Printf("\nNão foi possível mover dispositivo: O ambiente %s não existe!\n", ambOrigem)
			}

		case "9":
			fmt.Println("Destruindo ambientes vazios...")
			manager.RemoverAmbs(&espacoTupla)
		case "0":
			fmt.Println("TCHAU <3")
			break
		default:
			fmt.Println("Opção invalida!")
		}

		fmt.Printf("Pressione qualquer tecla + Enter para voltar ao menu principal...")
		var tecla string
		fmt.Scanf("%v", &tecla)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
