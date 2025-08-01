package logx

import "context"

// Non-contextual logging helpers
func Debug(msg string, fields map[string]interface{}) {
	Logger().Debug().Fields(fields).Msg(msg)
}

func Info(msg string, fields map[string]interface{}) {
	Logger().Info().Fields(fields).Msg(msg)
}

func Warn(msg string, fields map[string]interface{}) {
	Logger().Warn().Fields(fields).Msg(msg)
}

func Error(msg string, fields map[string]interface{}) {
	Logger().Error().Fields(fields).Msg(msg)
}

func Fatal(msg string, fields map[string]interface{}) {
	Logger().Fatal().Fields(fields).Msg(msg)
}

// Context-aware logging helpers
func DebugCtx(ctx context.Context, msg string, fields map[string]interface{}) {
	FromContext(ctx).Debug().Fields(fields).Msg(msg)
}

func InfoCtx(ctx context.Context, msg string, fields map[string]interface{}) {
	FromContext(ctx).Info().Fields(fields).Msg(msg)
}

func WarnCtx(ctx context.Context, msg string, fields map[string]interface{}) {
	FromContext(ctx).Warn().Fields(fields).Msg(msg)
}

func ErrorCtx(ctx context.Context, msg string, fields map[string]interface{}) {
	FromContext(ctx).Error().Fields(fields).Msg(msg)
}

func FatalCtx(ctx context.Context, msg string, fields map[string]interface{}) {
	FromContext(ctx).Fatal().Fields(fields).Msg(msg)
}
