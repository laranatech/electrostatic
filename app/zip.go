package app

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(source string, target string) error {
	newZipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	return filepath.Walk(source, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file == target {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name, err = filepath.Rel(filepath.Dir(source), file)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileReader, err := os.Open(file)
			if err != nil {
				return err
			}
			defer fileReader.Close()

			_, err = io.Copy(writer, fileReader)
			return err
		}

		return nil
	})
}
