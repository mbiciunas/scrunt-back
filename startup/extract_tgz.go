package startup

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExtractTarGz(path string, gzipStream io.Reader) {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			//pythonPath := filepath.Join(path, header.Name)

			//if err := os.Mkdir(header.Name, 0755); err != nil {
			if err := os.MkdirAll(filepath.Join(path, header.Name), 0755); err != nil {
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			//outFile, err := os.Create(header.Name)
			outFile, err := os.Create(filepath.Join(path, header.Name))
			if err != nil {
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:

			//
			// Lots to do to build the python tar file:
			// 		Compile python
			//		Resolve all link entries
			//		Remove unneeded files (bin?, etc)
			//
			// Need to automate setup of python so can be properly built.
			//
			log.Println("ExtractTarGz: uknown type: %s in %s", header.Typeflag, header.Name)
			//log.Fatalf("ExtractTarGz: uknown type: %s in %s",
			//	header.Typeflag,
			//	header.Name)
		}

	}
}
