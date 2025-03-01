/*
 * Copyright 2020 Matthew A. Titmus
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello net/http!\n"))
}

// Uncomment if you'll comment out gorillamux.go's main
/* func main() {
	// Associate a path with a handler function
	http.HandleFunc("/", helloGoHandler)

	// Bind to a port and default to the DefaultServeMux handler.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
