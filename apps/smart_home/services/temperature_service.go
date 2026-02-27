package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// TemperatureService handles fetching temperature data from external API
type TemperatureService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// PythonTemperatureResponse answer from python
type PythonTemperatureResponse struct {
	Temperature interface{} `json:"temperature"`
	Location    string      `json:"location"`
	SensorID    string      `json:"sensor_id"`
}

// TemperatureResponse represents the response from the temperature API
type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// NewTemperatureService creates a new temperature service
func NewTemperatureService(baseURL string) *TemperatureService {
	return &TemperatureService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// convertTemperature конвертирует температуру из разных типов в float64
func (s *TemperatureService) convertTemperature(temp interface{}) (float64, error) {
	switch v := temp.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case json.Number:
		return v.Float64()
	case string:
		var f float64
		_, err := fmt.Sscanf(v, "%f", &f)
		return f, err
	default:
		return 0, fmt.Errorf("unsupported temperature type: %T", v)
	}
}

// GetTemperature fetches temperature data for a specific location
func (s *TemperatureService) GetTemperature(location string) (*TemperatureResponse, error) {
	url := fmt.Sprintf("%s/temperature?location=%s", s.BaseURL, location)
	log.Printf("Calling temperature API: %s", url)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		log.Printf("Error calling temperature API: %v", err)
		return nil, fmt.Errorf("error fetching temperature data: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Temperature API response status: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var pythonResp PythonTemperatureResponse
	if err := json.NewDecoder(resp.Body).Decode(&pythonResp); err != nil {
		log.Printf("Error decoding temperature response: %v", err)
		return nil, fmt.Errorf("error decoding temperature response: %w", err)
	}

	log.Printf("Received from Python API - Temperature: '%v' (type: %T), Location: '%s', SensorID: '%s'",
		pythonResp.Temperature, pythonResp.Temperature, pythonResp.Location, pythonResp.SensorID)

	// Конвертируем температуру в float64
	value, err := s.convertTemperature(pythonResp.Temperature)
	if err != nil {
		log.Printf("Failed to convert temperature '%v': %v", pythonResp.Temperature, err)
		value = 0
	}

	// Определяем статус на основе температуры
	status := "normal"
	if value > 30 {
		status = "high"
	} else if value < 10 {
		status = "low"
	}

	log.Printf("Converted temperature value: %f, status: %s", value, status)

	return &TemperatureResponse{
		Value:       value,
		Unit:        "°C",
		Timestamp:   time.Now(),
		Location:    pythonResp.Location,
		Status:      status,
		SensorID:    pythonResp.SensorID,
		SensorType:  "temperature",
		Description: fmt.Sprintf("Temperature in %s is %.1f°C", pythonResp.Location, value),
	}, nil
}

// GetTemperatureByID fetches temperature data for a specific sensor ID
func (s *TemperatureService) GetTemperatureByID(sensorID string) (*TemperatureResponse, error) {
	url := fmt.Sprintf("%s/temperature?sensorID=%s", s.BaseURL, sensorID)
	log.Printf("Calling temperature API by ID: %s", url)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		log.Printf("Error calling temperature API: %v", err)
		return nil, fmt.Errorf("error fetching temperature data: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Temperature API response status: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var pythonResp PythonTemperatureResponse
	if err := json.NewDecoder(resp.Body).Decode(&pythonResp); err != nil {
		log.Printf("Error decoding temperature response: %v", err)
		return nil, fmt.Errorf("error decoding temperature response: %w", err)
	}

	log.Printf("Received from Python API - Temperature: '%v' (type: %T), Location: '%s', SensorID: '%s'",
		pythonResp.Temperature, pythonResp.Temperature, pythonResp.Location, pythonResp.SensorID)

	// Конвертируем температуру в float64
	value, err := s.convertTemperature(pythonResp.Temperature)
	if err != nil {
		log.Printf("Failed to convert temperature '%v': %v", pythonResp.Temperature, err)
		value = 0
	}

	// Определяем статус на основе температуры
	status := "normal"
	if value > 30 {
		status = "high"
	} else if value < 10 {
		status = "low"
	}

	log.Printf("Converted temperature value: %f, status: %s", value, status)

	return &TemperatureResponse{
		Value:       value,
		Unit:        "°C",
		Timestamp:   time.Now(),
		Location:    pythonResp.Location,
		Status:      status,
		SensorID:    pythonResp.SensorID,
		SensorType:  "temperature",
		Description: fmt.Sprintf("Temperature at sensor %s is %.1f°C", sensorID, value),
	}, nil
}