package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// FileUpload uploads local file to i2i.
func (c *Client) FileUpload(path string) (*FileRest, error) {

	var (
		filename = filepath.Base(path)
		buffer   bytes.Buffer
		writer   = multipart.NewWriter(&buffer)
	)

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err := part.Write(fileContent); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.nodeFileUploadAddress(), &buffer)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrHttpWithCode(resp.StatusCode)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fileModel := &FileRest{}

	if err := json.Unmarshal(respData, fileModel); err != nil {
		return nil, err
	}

	return fileModel, nil
}

// FileDownload downloads file from i2i to local host.
// Path argument is directory and resulting file is stored as path/{file_name}
// Note: if destination file already exist, error is returned.
func (c *Client) FileDownload(id string, path string) (*File, error) {
	file, err := c.file(id)
	if err != nil {
		return nil, err
	}
	localFilePath := filepath.Join(path, file.Name)

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, ErrNotADir()
	}

	if _, err := os.Stat(localFilePath); err == nil {
		return nil, ErrAlreadyExist
	}

	address := fmt.Sprintf("%s/%s", c.nodeFileDownloadAddress(), id)

	request, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrHttpWithCode(resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(localFilePath, data, 0644); err != nil {
		return nil, err
	}

	return file, nil
}

// FileRemove removes file with id from i2i.
func (c *Client) FileRemove(id string) (*File, error) {
	file := struct {
		File *File `json:"fileRemove"`
	}{}

	_, err := c.query(&query{
		query:     mutationFileRemove,
		timeout:   time.Second * 2,
		variables: map[string]interface{}{"fileID": id},
		response:  &file,
	})

	if err != nil {
		return nil, err
	}

	return file.File, nil
}

// File shows metadata of file with id.
func (c *Client) File(id string) (*File, error) {
	return c.file(id)
}

func (c *Client) file(id string) (*File, error) {
	file := struct {
		File *File `json:"file"`
	}{}

	_, err := c.query(&query{
		query:     queryFile,
		timeout:   time.Second * 2,
		variables: map[string]interface{}{"fileID": id},
		response:  &file,
	})

	if err != nil {
		return nil, err
	}

	return file.File, nil
}

// FileList lists files stored by the i2i
func (c *Client) FileList() ([]*File, error) {
	files := struct {
		Files []*File `json:"fileList"`
	}{}

	_, err := c.query(&query{
		query:     queryFileList,
		timeout:   time.Second * 2,
		variables: nil,
		response:  &files,
	})

	if err != nil {
		return nil, err
	}

	return files.Files, nil
}

// FileRename changes name of the file
func (c *Client) FileRename(id string, name string) (*File, error) {
	file := struct {
		File *File `json:"fileRename"`
	}{}

	_, err := c.query(&query{
		query:     mutationFileRename,
		timeout:   time.Second * 2,
		variables: map[string]interface{}{"fileID": id, "fileName": name},
		response:  &file,
	})

	if err != nil {
		return nil, err
	}

	return file.File, nil
}

// FileTransfer transfers file with id to connection.
// Connection is identified by its signature key.
func (c *Client) FileTransfer(id string, connection string) (*File, error) {
	file := struct {
		File *File `json:"fileTransfer"`
	}{}

	_, err := c.query(&query{
		query:     mutationFileTransfer,
		timeout:   time.Second * 2,
		variables: map[string]interface{}{"fileID": id, "connectionKey": connection},
		response:  &file,
	})

	if err != nil {
		return nil, err
	}

	return file.File, nil
}
