package initializers

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)

func Supabase() *supa.Client {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	return supabase
}
