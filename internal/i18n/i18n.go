package i18n

import (
    "encoding/json"
    "io/ioutil"
    "sync"
)

var (
    mu      sync.Mutex
    messages = make(map[string]map[string]string)
)

func LoadTranslations() error {
    mu.Lock()
    defer mu.Unlock()

    files, err := ioutil.ReadDir("locales")
    if err != nil {
        return err
    }

    for _, file := range files {
        if file.IsDir() || file.Name() == ".DS_Store" {
            continue
        }

        data, err := ioutil.ReadFile("locales/" + file.Name())
        if err != nil {
            return err
        }

        var translations map[string]string
        if err := json.Unmarshal(data, &translations); err != nil {
            return err
        }

        lang := file.Name()[:len(file.Name())-5] // Remove .json
        messages[lang] = translations
    }

    return nil
}

func Translate(lang, key string) string {
    mu.Lock()
    defer mu.Unlock()

    if translation, ok := messages[lang][key]; ok {
        return translation
    }
    return key // Fallback to the key itself if no translation is found
}