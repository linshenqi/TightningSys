package cvi3

import (
	"fmt"
	"strconv"
	"time"
	"strings"
	"encoding/xml"
  "server"
)

type CVI3 struct {
  server  *CVI3Server
}
