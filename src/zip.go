
package main

import (
  "archive/zip"
  "io"
  "os"
  "path/filepath"
  "strings"
  "path"
)

func zip_folder(ch *Chapter) error {
  source := path.Join(ch.Manga, ch.Chapter)
  target := path.Join(ch.Manga, "archive", ch.Manga + "_" + ch.Chapter + ".cbz")

  file, err := os.Create(target)
  if err != nil { return err }
  defer file.Close()

  zip_writer := zip.NewWriter(file)
  defer zip_writer.Close()

  base_folder := filepath.Base(source)

  return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
    if err != nil { return err }
    if path == source { return nil }

    header, err := zip.FileInfoHeader(info)
    if err != nil { return err }
    header.Name = filepath.Join(base_folder, strings.TrimPrefix(path, source))

    if info.IsDir() {
      header.Name += "/"
      _, err = zip_writer.CreateHeader(header)
      return err
    }

    file, err := os.Open(path)
    if err != nil { return err }
    defer file.Close()

    writer, err := zip_writer.CreateHeader(header)
    if err != nil {
      return err
    }

    _, err = io.Copy(writer, file)
    return err
  })
}
