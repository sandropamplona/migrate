package UtilsDSIN

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/nakagami/firebirdsql"

	iurl "github.com/golang-migrate/migrate/v4/internal/url"
)

func sliceURL(url string) (string, string) {
	driveName, _ := iurl.SchemeFromURL(url)

	indice := strings.Index(url, ":")

	loginPathDB := url[indice+3 : len(url)]
	return driveName, loginPathDB
}
func IdentificadorDB(urlDatabase string) (string, error) {
	var resultSet string

	driverName, PathBanco := sliceURL(urlDatabase)
	driverName = fmt.Sprintf("%s%s", driverName, "sql")

	conn, _ := sql.Open(driverName, PathBanco)
	defer conn.Close()

	if err := conn.QueryRow("SELECT TRIM('SGITDADOS') AS BANCO FROM RDB$RELATIONS WHERE RDB$RELATIONS.RDB$SYSTEM_FLAG =0 AND  RDB$RELATIONS.RDB$RELATION_NAME='ACARQUIVOB' UNION SELECT TRIM('SGITIMAGENS') FROM RDB$RELATIONS WHERE RDB$RELATIONS.RDB$SYSTEM_FLAG =0 AND RDB$RELATIONS.RDB$RELATION_NAME='LCFOTOINFRACAO'").Scan(&resultSet); err != nil {
		if err != nil {
			return resultSet, err
		}
	}

	return resultSet, nil
}
