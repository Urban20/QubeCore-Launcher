package data

// este modulo tiene las estructuras que se necesitan para parsear el json de manifiest.json
// para cada version

type Artifact struct { // estructura artifact se usa para si mismo y para classifiers
	Path string `json:"path"`
	SHA1 string `json:"sha1"`
	Size int64  `json:"size"`
	URL  string `json:"url"`
}

type Library struct {
	Downloads struct {
		Artifact    Artifact            `json:"artifact"`
		Classifiers map[string]Artifact `json:"classifiers"`
	} `json:"downloads"`
	Natives map[string]string `json:"natives"`
	Rules   []struct {
		Action string `json:"action"`
		OS     struct {
			Name string `json:"name"`
		} `json:"os"`
	} `json:"rules"`
}

type VersionJSON struct {
	ID         string `json:"id"`
	MainClass  string `json:"mainClass"`
	AssetIndex struct {
		ID   string `json:"id"`
		SHA1 string `json:"sha1"`
		URL  string `json:"url"`
	} `json:"assetIndex"`
	Downloads struct {
		Client Artifact `json:"client"`
	} `json:"downloads"`
	Libraries []Library `json:"libraries"`
}

type AssetIndex struct {
	Objects map[string]struct {
		Hash string `json:"hash"`
		Size int64  `json:"size"`
	} `json:"objects"`
}

type Task struct {
	URL      string // url
	DestPath string //directorio destino
	SHA1     string //hash
	Label    string //etiqueta
}
