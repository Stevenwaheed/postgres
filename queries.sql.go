// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package postgers

import (
	"context"
	"database/sql"
)

const newUserTransaction = `-- name: newUserTransaction :exec
    INSERT INTO users("name", "phone_number") VALUES ($1, $2)
`

type newUserTransactionParams struct {
	Name        sql.NullString `json:"name"`
	PhoneNumber string         `json:"phone_number"`
}

func (q *Queries) newUserTransaction(ctx context.Context, arg newUserTransactionParams) error {
	_, err := q.db.ExecContext(ctx, newUserTransaction, arg.Name, arg.PhoneNumber)
	return err
}

const otpTransaction = `-- name: otpTransaction :exec
    INSERT INTO users_otp("otp", "otp_expiration_time", "phone_number_otp") VALUES ($1, $2, $3)
`

type otpTransactionParams struct {
	Otp               sql.NullString `json:"otp"`
	OtpExpirationTime sql.NullTime   `json:"otp_expiration_time"`
	PhoneNumberOtp    string         `json:"phone_number_otp"`
}

func (q *Queries) otpTransaction(ctx context.Context, arg otpTransactionParams) error {
	_, err := q.db.ExecContext(ctx, otpTransaction, arg.Otp, arg.OtpExpirationTime, arg.PhoneNumberOtp)
	return err
}

const verifyOTP = `-- name: verifyOTP :one
    SELECT otp, otp_expiration_time FROM users_otp WHERE phone_number_otp=$1 AND otp=$2
`

type verifyOTPParams struct {
	PhoneNumberOtp string         `json:"phone_number_otp"`
	Otp            sql.NullString `json:"otp"`
}

type verifyOTPRow struct {
	Otp               sql.NullString `json:"otp"`
	OtpExpirationTime sql.NullTime   `json:"otp_expiration_time"`
}

func (q *Queries) verifyOTP(ctx context.Context, arg verifyOTPParams) (verifyOTPRow, error) {
	row := q.db.QueryRowContext(ctx, verifyOTP, arg.PhoneNumberOtp, arg.Otp)
	var i verifyOTPRow
	err := row.Scan(&i.Otp, &i.OtpExpirationTime)
	return i, err
}
