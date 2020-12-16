package exporter

import (
	"encoding/csv"
	"io/ioutil"
	"os"
)

// Csv 导出助手
type Exporter struct {
	file   *os.File
	data   [][]string
	writer *csv.Writer
}

// 新建导出助手
// title 头部标题
func NewExporter(title []string) (*Exporter, error) {
	file, err := ioutil.TempFile("", "csv")
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString("\xEF\xBB\xBF")
	if err != nil {
		return nil, err
	}
	return &Exporter{
		file:   file,
		writer: csv.NewWriter(file),
		data:   [][]string{title},
	}, nil
}

// 关闭并删除文件
func (c *Exporter) Close() error {
	if err := c.file.Close(); err != nil {
		return err
	}
	return os.Remove(c.file.Name())
}

// 追加数据
func (c *Exporter) Append(data []string) {
	c.data = append(c.data, data)
}

// 写入Csv并返回文件
func (c *Exporter) WriteAllSeekZero() (*os.File, error) {
	err := c.writeAll()
	if err != nil {
		return nil, err
	}
	return c.seekZero()
}

// 写入
func (c *Exporter) writeAll() error {
	err := c.writer.WriteAll(c.data)
	if err != nil {
		return err
	}
	c.writer.Flush()
	return nil
}

func (c *Exporter) seekZero() (*os.File, error) {
	if _, err := c.file.Seek(0, 0); err != nil {
		return nil, err
	}
	return c.file, nil
}
