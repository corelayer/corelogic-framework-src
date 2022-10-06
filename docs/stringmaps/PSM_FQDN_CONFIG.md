# PSM_FQDN_CONFIG

This stringmap allows the configuration of all FQDN in the configuration
## Format
| Key                             | Value                             |
|:--------------------------------|:----------------------------------|
| [fqdn](###fqdn)=\<domain>;      | [acme](###acme)=\<acme_value>;    |

## Keys
### fqdn
Specify the fqdn for this entry

## Values
### acme
Speficy if the ACME protocol is enabled for this fqdn
#### Allowed values
- enabled
- disabled