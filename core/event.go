package core

import "app/core/http"

type Listener map[string]func(*http.Event)
