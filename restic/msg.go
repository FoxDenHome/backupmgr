package restic

import (
	"encoding/json"
	"log"
	"time"
)

type Message struct {
	MessageType string `json:"message_type"`
}

type Summary struct {
	Message

	FilesNew            int       `json:"files_new"`
	FilesChanged        int       `json:"files_changed"`
	FilesUnmodified     int       `json:"files_unmodified"`
	DirsNew             int       `json:"dirs_new"`
	DirsChanged         int       `json:"dirs_changed"`
	DirsUnmodified      int       `json:"dirs_unmodified"`
	DaraBlobs           int       `json:"data_blobs"`
	TreeBlobs           int       `json:"tree_blobs"`
	DataAdded           int       `json:"data_added"`
	DataAddedPacked     int       `json:"data_added_packed"`
	TotalFilesProcessed int       `json:"total_files_processed"`
	TotalBytesProcessed int       `json:"total_bytes_processed"`
	TotalDuration       float64   `json:"total_duration"`
	BackupStart         time.Time `json:"backup_start"`
	BackupEnd           time.Time `json:"backup_end"`
	SnapshotID          string    `json:"snapshot_id"`
}

type Status struct {
	Message

	SecondsElapsed int     `json:"seconds_elapsed"`
	PercentDone    float64 `json:"percent_done"`
	TotalFiles     int     `json:"total_files"`
	FilesDone      int     `json:"files_done"`
	TotalBytes     int     `json:"total_bytes"`
	BytesDone      int     `json:"bytes_done"`
}

type ExitError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *Repo) handleNextMessage(decoder *json.Decoder) error {
	var rawMsg json.RawMessage

	if err := decoder.Decode(&rawMsg); err != nil {
		return err
	}

	var baseMsg Message
	if err := json.Unmarshal(rawMsg, &baseMsg); err != nil {
		return err
	}

	switch baseMsg.MessageType {
	case "summary":
		var summary Summary
		if err := json.Unmarshal(rawMsg, &summary); err != nil {
			return err
		}
		log.Printf("[RESTIC] Summary: %+v", summary)
	case "status":
		var status Status
		if err := json.Unmarshal(rawMsg, &status); err != nil {
			return err
		}
		log.Printf("[RESTIC] Status: %+v", status)
	case "exit_error":
		var exitErr ExitError
		if err := json.Unmarshal(rawMsg, &exitErr); err != nil {
			return err
		}
		log.Fatalf("[RESTIC] Fatal error: %+v", exitErr)
	default:
		// Ignore other message types for now
	}

	return nil
}
