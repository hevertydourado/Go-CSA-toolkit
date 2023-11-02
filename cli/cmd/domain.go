package cmd

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"

	"github.com/fatih/color"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"github.com/spf13/cobra"
	"golang.org/x/net/publicsuffix"
)

var domain string

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Get information about a domain",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a domain as an argument.")
			return
		}

		domain = args[0]
		domainRetrieveInfo()
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)
}

func hasSubdomain(domain string) bool {
	u, err := url.Parse("https://" + domain)
	if err != nil {
		log.Printf("Invalid domain: %s\n", domain)
		return false
	}

	hostParts := strings.Split(u.Hostname(), ".")
	tld, icann := publicsuffix.PublicSuffix(u.Hostname())

	// Check if it's not a recognized TLD or has subdomains
	return tld != "" && tld != u.Hostname() && icann && len(hostParts) > 2
}

func domainRetrieveInfo() {
	atention := color.New(color.FgYellow, color.Bold).Sprint("[ ! ]")
	warning := color.New(color.FgRed, color.Bold).Sprint("[ X ]")
	success := color.New(color.FgGreen, color.Bold).Sprint("[ ✅ ]")

	fmt.Printf("%s Tip: Try to search without a subdomain: ex.: youtube.com instead of www.youtube.com... or the other way around...\n\n", atention)

	fmt.Printf("%s Retrieving domain information...\n", atention)

	hasSubdomain := hasSubdomain(domain)

	// IP A and AAA records
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Printf("Could not obtain IP for %s. Please check the domain and try again.\n\n", domain)
	} else {
		if len(ips) > 0 {
			fmt.Printf("Domain Information for: %s\n\n", domain)
			fmt.Println("IPs: (A and AAA records)")
			for _, ip := range ips {
				fmt.Printf("%s IP: %s\n", success, ip)
			}
		} else {
			fmt.Printf("%s No IP information available for domain: %s\n", warning, domain)
		}
	}

	if hasSubdomain {
		// CNAME records
		cname, err := net.LookupCNAME(domain)
		if err != nil {
			log.Printf("%s Could not obtain CNAME records for %s. Please check the domain and try again.\n\n", warning, domain)
		} else {
			fmt.Println("\nCNAME records:")
			fmt.Printf("%s %s\n", success, cname)
		}
	}

	// Who.is querying
	if !hasSubdomain {
		// MX records
		mxs, err := net.LookupMX(domain)
		if err != nil {
			log.Printf("%s Could not obtain MX records for %s. Please check the domain and try again.\n\n", warning, domain)
		} else {
			if len(mxs) > 0 {
				fmt.Println("\nMX records: ")
				for _, mx := range mxs {
					fmt.Printf("%s MX: %s\n", success, mx.Host)
				}
			} else {
				fmt.Printf("%s No MX records information available for domain: %s\n", warning, domain)
			}
		}

		// TXT records
		txts, err := net.LookupTXT(domain)
		if err != nil {
			log.Printf("%s Could not obtain TXT records for %s. Please check the domain and try again.\n\n", warning, domain)
		} else {
			if len(txts) > 0 {
				fmt.Println("\nTXT records: ")
				for _, txt := range txts {
					fmt.Printf("%s TXT: %s\n", success, txt)
				}
			} else {
				fmt.Printf("%s No TXT records information available for domain: %s\n", warning, domain)
			}
		}

		// NS records
		nss, err := net.LookupNS(domain)
		if err != nil {
			log.Printf("Could not obtain NS records for %s. Please check the domain and try again.\n\n", domain)
		} else {
			if len(nss) > 0 {
				fmt.Printf("Domain Information for: %s\n\n", domain)
				fmt.Println("NS records:")
				for _, ns := range nss {
					fmt.Printf("%s NS: %s\n", success, ns.Host)
				}
			} else {
				fmt.Printf("%s No NS records available for domain: %s\n", warning, domain)
			}
		}

		fmt.Println("\nWho.is querying...")
		result, err := whois.Whois(domain)
		if err != nil {
			log.Printf("%s Could not obtain Who.is information for %s. Please check the domain and try again.\n\n", warning, domain)
			return
		}
		parsedResult, err := whoisparser.Parse(result)
		if err != nil {
			log.Println(err)
			return
		}

		status := parsedResult.Domain.Status
		creationDate := parsedResult.Domain.CreatedDate
		expirationDate := parsedResult.Domain.ExpirationDate
		registrantName := parsedResult.Registrant.Name
		registrantMail := parsedResult.Registrant.Email
		ownerId := parsedResult.Domain.ID
		name := parsedResult.Domain.Name

		fmt.Printf("%s Domain name: %s\n", success, name)
		fmt.Printf("%s Domain status: %s\n", success, status)
		fmt.Printf("%s Domain Owner ID: %s\n", success, ownerId)
		fmt.Printf("%s Domain creation date: %s [MM/DD/YYYY]\n", success, creationDate)
		fmt.Printf("%s Domain expiration date: %s [MM/DD/YYYY]\n", success, expirationDate)
		fmt.Printf("%s Domain registrant name: %v\n", success, registrantName)
		fmt.Printf("%s Domain registrant Email: %s\n", success, registrantMail)
	}
}
