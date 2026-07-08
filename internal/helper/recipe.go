package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type AIGeneratedRecipe struct {
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	CookingTime        int      `json:"cooking_time"`
	IngredientsUsed    string   `json:"ingredients_used"`
	MissingIngredients string   `json:"missing_ingredients"`
	Instructions       []string `json:"instructions"`
}

func GenerateRecipesFromAI(foodNames []string) ([]*AIGeneratedRecipe, error) {
	apiKey := os.Getenv("OPEN_ROUTER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPEN_ROUTER_API_KEY not set")
	}

	url := "https://openrouter.ai/api/v1/chat/completions"

	prompt := fmt.Sprintf(
		`Kamu adalah koki profesional. Berikut adalah bahan-bahan makanan yang tersedia di kulkas:
%s

Buatlah 5 resep masakan yang bisa dibuat dari bahan-bahan tersebut. Setiap resep boleh menggunakan bahan tambahan lain yang tidak ada dalam daftar jika diperlukan.

Untuk setiap resep, berikan dalam format JSON dengan field:
- title: string (judul resep)
- description: string (deskripsi singkat)
- cooking_time: number (waktu memasak dalam menit)
- ingredients_used: string (bahan-bahan dari daftar yang digunakan, dipisah koma)
- missing_ingredients: string (bahan tambahan yang diperlukan tapi tidak ada di daftar, dipisah koma. Jika tidak ada, isi string kosong)
- instructions: array of strings (langkah-langkah memasak)

Output berupa array JSON dengan tepat 5 resep. Hanya output JSON, tanpa teks lain.`,
		strings.Join(foodNames, ", "),
	)

	reqBody := map[string]interface{}{
		"model": "openai/gpt-4o-mini",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %s", string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}

	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	content := message["content"].(string)

	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var recipes []*AIGeneratedRecipe
	if err := json.Unmarshal([]byte(content), &recipes); err != nil {
		return nil, fmt.Errorf("failed to parse recipes: %s\nraw: %s", err.Error(), content)
	}

	return recipes, nil
}
