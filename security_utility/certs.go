package security_utility

import (
    "crypto/x509"
)

/*
    Loads an arbitrary number of certitificates from passed string paths
*/
func loadCertPool(cert_paths ...string) (cert_pool *x509.CertPool, err error)  {
    cert_pool = x509.NewCertPool()
    for _, cert_path := range cert_paths {
        cert_pem, err := ioutil.ReadFile(cert_path)
        if err != nil {
            log.Print("Read Cert PEM: ", err)
            return cert_pool, err
        }
        cert_pool.AppendCertsFromPEM(cert_pem)
    }
    return cert_pool, nil
}
