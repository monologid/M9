package login

import (
	"strings"

	"github.com/monologid/m9/instrument"
	"github.com/monologid/m9/serviceprovider"
)

var prometheusModuleName = "login"

var (
	PrometheusDefaultTotal                = instrument.NewPrometheus().NewCounter(prometheusModuleName, "default_total", "The total number of login using Default.")
	PrometheusDefaultSuccessTotal         = instrument.NewPrometheus().NewCounter(prometheusModuleName, "default_success_total", "The total number of success login using Default.")
	PrometheusDefaultFailedTotal          = instrument.NewPrometheus().NewCounter(prometheusModuleName, "default_failed_total", "The total number of failed login using Default.")
	PrometheusRegisterDefaultTotal        = instrument.NewPrometheus().NewCounter(prometheusModuleName, "default_register_total", "The total number of account registration using Default.")
	PrometheusRegisterDefaultSuccessTotal = instrument.NewPrometheus().NewCounter(prometheusModuleName, "default_register_success_total", "The total number of success account registration using Default.")
	PrometheusRegisterDefaultFailedTotal  = instrument.NewPrometheus().NewCounter(prometheusModuleName, "default_register_failed_total", "The total number of failed account registration using Default.")

	PrometheusFacebookTotal                = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_total", "The total number of login using Facebook.")
	PrometheusFacebookSuccessTotal         = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_success_total", "The total number of success login using Facebook.")
	PrometheusFacebookFailedTotal          = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_failed_total", "The total number of failed login using Facebook.")
	PrometheusRegisterFacebookTotal        = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_register_total", "The total number of account registration using Facebook.")
	PrometheusRegisterFacebookSuccessTotal = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_register_success_total", "The total number of success account registration using Facebook.")
	PrometheusRegisterFacebookFailedTotal  = instrument.NewPrometheus().NewCounter(prometheusModuleName, "facebook_register_failed_total", "The total number of failed account registration using Facebook.")

	PrometheusGoogleTotal                = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_total", "The total number of login using Google.")
	PrometheusGoogleSuccessTotal         = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_success_total", "The total number of success login using Google.")
	PrometheusGoogleFailedTotal          = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_failed_total", "The total number of failed login using Google.")
	PrometheusRegisterGoogleTotal        = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_register_total", "The total number of account registration using Google.")
	PrometheusRegisterGoogleSuccessTotal = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_register_success_total", "The total number of success account registration using Google.")
	PrometheusRegisterGoogleFailedTotal  = instrument.NewPrometheus().NewCounter(prometheusModuleName, "google_register_failed_total", "The total number of failed account registration using Google.")
)

// MetricInitiateLogin increases the metric when iniate a login
func MetricInitiateLogin(provider string) {
	switch strings.ToUpper(provider) {
	case serviceprovider.FACEBOOK:
		PrometheusFacebookTotal.Inc()
	case serviceprovider.GOOGLE:
		PrometheusGoogleTotal.Inc()
	default:
		PrometheusDefaultTotal.Inc()
	}
}

// MetricLoginSuccess increases the metric when login is success
func MetricLoginSuccess(provider string) {
	switch strings.ToUpper(provider) {
	case serviceprovider.FACEBOOK:
		PrometheusFacebookSuccessTotal.Inc()
	case serviceprovider.GOOGLE:
		PrometheusGoogleSuccessTotal.Inc()
	default:
		PrometheusDefaultSuccessTotal.Inc()
	}
}

// MetricLoginFailed increases the metric when login is failed
func MetricLoginFailed(provider string) {
	switch strings.ToUpper(provider) {
	case serviceprovider.FACEBOOK:
		PrometheusFacebookFailedTotal.Inc()
	case serviceprovider.GOOGLE:
		PrometheusGoogleFailedTotal.Inc()
	default:
		PrometheusDefaultFailedTotal.Inc()
	}
}

// MetricInitiateAccountRegistration increases the metric when there's a new account registration
func MetricInitiateAccountRegistration(provider string) {
	switch strings.ToUpper(provider) {
	case serviceprovider.FACEBOOK:
		PrometheusRegisterFacebookTotal.Inc()
	case serviceprovider.GOOGLE:
		PrometheusRegisterGoogleTotal.Inc()
	default:
		PrometheusRegisterDefaultTotal.Inc()
	}
}

// MetricInitiateAccountRegistrationSuccess increases the metric when new account registration is success
func MetricInitiateAccountRegistrationSuccess(provider string) {
	switch strings.ToUpper(provider) {
	case serviceprovider.FACEBOOK:
		PrometheusRegisterFacebookSuccessTotal.Inc()
	case serviceprovider.GOOGLE:
		PrometheusRegisterGoogleSuccessTotal.Inc()
	default:
		PrometheusRegisterDefaultSuccessTotal.Inc()
	}
}

// MetricInitiateAccountRegistrationFailed increases the metric when new account registration is success
func MetricInitiateAccountRegistrationFailed(provider string) {
	switch strings.ToUpper(provider) {
	case serviceprovider.FACEBOOK:
		PrometheusRegisterFacebookFailedTotal.Inc()
	case serviceprovider.GOOGLE:
		PrometheusRegisterGoogleFailedTotal.Inc()
	default:
		PrometheusRegisterDefaultFailedTotal.Inc()
	}
}
