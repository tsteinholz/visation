// apis/shadertoy.go -- Visation
// Copyright (C) 2020 Thomas Steinholz
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Provides an interface to ShaderToy's open APIs.
// Docstrings retrieved from: https://www.shadertoy.com/howto

package shadertoy

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// All API calls require a key that you can request (check the first section of
// this page) for free and returns JSON files that you can easily read with your
// favorite parser.
type ShaderToy() struct {
    api_key string
}

// Builds a new shadertoy instance, requires the API key.
func NewShaderToy(api_key string) *ShaderToy {
    return &ShaderToy{api_key: api_key}
}

// Provides the final URL to call the API with. Requires REST call and params as
// the ext (extention) argument.
func (st ShaderToy) BuildURL(ext string, params string) string {
    const base_url = "https://www.shadertoy.com/api/v1/shaders/"
    return base_url + ext + "?key=" st.api_key + "/" + params
}

// Get a shader from a shader ID.
//
// where shaderID is the same ID used in the Shadertoy URLs, and
// also the values returned by the "Query Shaders".
func (st ShaderToy) SearchShader(field string) error {
    // https://www.shadertoy.com/api/v1/shaders/shaderID?key=appkey
    url := st.BuildURL(shader_id)
    response, err := http.Get(url)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
        return err
    }
    data, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(data))
    return nil
}

// Get all shaders.
//
// https://www.shadertoy.com/api/v1/shaders?key=appkey
func (st SharderToy) GetShaders() {

}

// Access the assets.
//
// When you retrieve a shader you will see a key called "inputs", this can be
// a texture/video/keyboard/sound used by the shader. The JSON returned when
// accessing a shader will look like this:
//
//  [..]{"inputs":[{"id":17,
//                  "src":"/media/a/(hash.extension)",
//                  "ctype":"texture",
//                  "channel":0
//                 }[..]
//
// To access this specific asset you can just cut and paste this path
// https://www.shadertoy.com/media/a/(hash.extension)
func GetAsset() {

}

// Query shaders sorted by "name", "love", "popular", "newest", "hot" (by
// default, it uses "popular").
//
// https://www.shadertoy.com/api/v1/shaders/query/string?sort=newest&key=appkey
//
//
// Query shaders with paging. Define a "from" and a "num" of shaders that
// you want (by default, there is no paging)
//
// https://www.shadertoy.com/api/v1/shaders/query/string?from=5&num=25&key=appkey
//
// Query shaders with filters: "vr", "soundoutput", "soundinput", "webcam",
// "multipass", "musicstream" (by default, there is no filter)
//
// https://www.shadertoy.com/api/v1/shaders/query/string?filter=vr&key=appkey
func QueryShaders(sorted string, filters string, author string, num int) {

}


func main() {
    fmt.Println("Starting the application...")
    response, err := http.Get("https://httpbin.org/ip")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData)
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    fmt.Println("Terminating the application...")
}
