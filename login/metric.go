package login

import (
	"github.com/monologid/m9/instrument"
)

var prometheusModuleName = "login"

var (
	PrometheusLoginDefaultTotal         = instrument.NewPrometheus().NewCounter(prometheusModuleName, "default_total", "The total number of login using Default.")
	PrometheusLoginFacebookTotal        = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_total", "The total number of login using Facebook.")
	PrometheusLoginFacebookSuccessTotal = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_success_total", "The total number of success login using Facebook.")
	PrometheusLoginGoogleTotal          = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_total", "The total number of login using Google.")
	PrometheusLoginGoogleSuccessTotal   = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_success_total", "The total number of success login using Google.")
	PrometheusLoginTwitterTotal         = instrument.NewPrometheus().NewCounter(prometheusModuleName, "twitter_total", "The total number of login using Twitter.")
	PrometheusLoginTwitterSuccessTotal  = instrument.NewPrometheus().NewCounter(prometheusModuleName, "twitter_success_total", "The total number of success login using Twitter.")
)
