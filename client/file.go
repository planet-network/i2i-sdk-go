package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (c *Client) FileUpload(path string) (*File, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.nodeFileUploadAddress(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	_, fileName := filepath.Split(path)
	request.Header.Set("filename", fileName)

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

	fileModel := &File{}

	if err := json.Unmarshal(respData, fileModel); err != nil {
		return nil, err
	}

	return fileModel, nil
}

func (c *Client) FileDownload(id string, path string) (*File, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, ErrNotADir()
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

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) FileRemove(id string) (string, error) {
	file := struct {
		ID string `json:"fileRemove"`
	}{}

	_, err := c.query(&query{
		query:     mutationFileRemove,
		timeout:   time.Second * 2,
		variables: map[string]interface{}{"fileID": id},
		response:  &file,
	})

	if err != nil {
		return "", err
	}

	return file.ID, nil
}

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
