package gencore

import (
	"context"
	"errors"
	"fmt"
	"github.com/52lu/go-helpers/logutil"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type dbLogger struct {
	dbConf gormlogger.Config
}

var _dbLogger = new(dbLogger)

func NewDBLogger(dbConf gormlogger.Config) *dbLogger {
	return &dbLogger{
		dbConf: dbConf,
	}
}

func (d *dbLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return _dbLogger
}

func (d *dbLogger) Info(ctx context.Context, s string, i ...interface{}) {
	logutil.Infof(ctx, s, i...)
}

func (d *dbLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	logutil.Warnf(ctx, s, i...)
}

func (d *dbLogger) Error(ctx context.Context, s string, i ...interface{}) {
	logutil.Errorf(ctx, s, i...)
}

func (d *dbLogger) Printf(ctx context.Context, s string, i ...interface{}) {
	logutil.Infof(ctx, s, i...)
}

func (d *dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if d.dbConf.LogLevel <= gormlogger.Silent {
		return
	}
	traceInfoFormat := "[%.3fms] [rows:%v] %s"
	traceFormat := "%s [%.3fms] [rows:%v] %s"
	elapsed := time.Since(begin)
	switch {
	case err != nil && d.dbConf.LogLevel >= gormlogger.Error && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !d.dbConf.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			d.Error(ctx, traceFormat, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			d.Error(ctx, traceFormat, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > d.dbConf.SlowThreshold && d.dbConf.SlowThreshold != 0 && d.dbConf.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", d.dbConf.SlowThreshold)
		if rows == -1 {
			d.Printf(ctx, traceFormat, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			d.Printf(ctx, traceFormat, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case d.dbConf.LogLevel == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {
			d.Printf(ctx, traceInfoFormat, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			d.Printf(ctx, traceInfoFormat, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
