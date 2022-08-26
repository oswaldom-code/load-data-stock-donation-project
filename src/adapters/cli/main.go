package cli

import (
	"fmt"

	"github.com/oswaldom-code/load-data-stock-donation-project/src/services"
	"github.com/spf13/cobra"
)

const (
	CONFIRMATION = "> Enter Yes to confirm or No to cancel (Y/n): "
)

func askConfirmation() bool {
	var answer string
	// ask for confirmation
	fmt.Print(CONFIRMATION)
	fmt.Scanln(&answer)
	if answer == "Y" || answer == "y" {
		return true
	}
	return false
}

func RunCliCmd(cmd *cobra.Command, args []string) error {
	service := services.NewLoadDataServices()
	function, err := cmd.Flags().GetString("function")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	switch function {
	case "load":
		target, err := cmd.Flags().GetString("target")
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		if target == "file" {
			path, err := cmd.Flags().GetString("path")
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
			if askConfirmation() {
				result, err := service.LoadDataFromJsonFile(path)
				if err != nil {
					fmt.Println("Error:", err)
					return err
				}
				fmt.Println("Inserted", result, "rows")

			}
		} else if target == "directory" {
			path, err := cmd.Flags().GetString("path")
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
			if askConfirmation() {
				result, err := service.LoadDataFromDirectory(path)
				if err != nil {
					fmt.Println("Error:", err)
					return err
				}
				fmt.Println("Inserted", result, "rows")
			}
		} else {
			fmt.Println("Error:", err)
			return err
		}

	case "test":
		err = service.TestDbConnection()
		if err != nil {
			fmt.Printf("> ❌Test error: %s\n", err)
			return err
		}
		fmt.Println("> ✅ Test connection to database success ")
		return nil
	default:
		fmt.Println("> ❌ Error: Invalid function")
	}
	return nil
}
