package startup

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func extractTarGz(path string, gzipStream io.Reader) {
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
			err = outFile.Close()
			if err != nil {
				log.Fatalf("ExtractTarGz: Close() failed: %s", err.Error())
			}

		default:
			log.Printf("ExtractTarGz: unknown type: %c in %s\n", header.Typeflag, header.Name)
			//log.Fatalf("ExtractTarGz: unknown type: %s in %s",
			//	header.Typeflag,
			//	header.Name)
		}

	}
}
