package controllers

import (
	"bufio"
	_ "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/internal/service/api/pkg/resp"
)

func GetLogs(c *fiber.Ctx) error {
	// Open the logs.log file
	file, err := os.Open("log.log")
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check the number of lines
	if len(lines) > 1000 {
		// If more than 1000 lines, return the last 200 lines
		lines = lines[len(lines)-200:]
	}

	// Get the last modification time of the file
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	lastModified := fileInfo.ModTime()

	// 返回响应
	result := models.LogResult{
		Lines:          lines,
		LastModifiedAt: lastModified.String(),
	}

	res := resp.JsonResp(result)

	return c.Status(fiber.StatusOK).JSON(res)
}
