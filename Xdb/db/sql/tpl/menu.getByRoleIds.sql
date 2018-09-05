SELECT
    m1.id AS id ,
    m1.icon AS icon ,
    (
        CASE
        WHEN(m2.id = 0 OR m2.id IS NULL) THEN
            0
        ELSE
            m2.id
        END
    ) AS pcode ,
    m1.name AS name,
    m1.url AS url,
    m1.levels AS levels ,
    m1.ismenu AS ismenu ,
    m1.num AS num
FROM
    sys_menu m1
LEFT JOIN sys_menu m2 ON m1.pcode = m2.code
INNER JOIN(
    SELECT
        ID
    FROM
        sys_menu
    WHERE
        ID IN(
            SELECT
                menuid
            FROM
                sys_relation rela
            WHERE
                rela.roleid in (
            {{ range $index, $value := .roleIds }}
                {{ if ne $index $.length }}
                    {{$value}},
                {{ else }}
                    {{$value}}
                {{ end }}
            {{ end }}
                )
        )
) m3 ON m1.id = m3.id
WHERE
    m1.ismenu = 1 AND m1.status = 1
ORDER BY
    levels ,
    num ASC	