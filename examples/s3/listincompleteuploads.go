// +build ignore

/*
 * Minio Go Library for Amazon S3 Compatible Cloud Storage (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"

	"github.com/minio/minio-go"
)

func main() {
	// Requests are always secure by default. set inSecure=true to enable insecure access.
	// inSecure boolean is the last argument for New().

	// New provides a client object backend by automatically detected signature type based
	// on the provider.
	s3Client, err := minio.New("s3.amazonaws.com", "YOUR-ACCESS-KEY-HERE", "YOUR-SECRET-KEY-HERE", false)
	if err != nil {
		log.Fatalln(err)
	}

	for multipartObject := range s3Client.ListIncompleteUploads("bucket-name", "objectName", true) {
		if multipartObject.Err != nil {
			log.Fatalln(multipartObject.Err)
		}
		log.Println(multipartObject)
	}
}
