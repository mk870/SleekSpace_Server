package supabase

import (
	"SleekSpace/utilities"

	storage_go "github.com/supabase-community/storage-go"
)

func StorageConnect() *storage_go.Client {
	storageClient := storage_go.NewClient("https://"+utilities.GetEnvVariables().SupabaseReferenceId+".supabase.co/storage/v1", utilities.GetEnvVariables().SupabaseApiKey, nil)
	return storageClient
}
