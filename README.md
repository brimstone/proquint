proquit
=======

Simple tool to encode or decode stdin or the only flagless argument to and from
proquit.

Example
-------

### Plain
```
$ proquint 'Hello!'
jodoh-soduk-dosuk
```

```
$ proquint jodoh-soduk-dosuk
Hello!
```

### Base58
```
$ proquint -base58 QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB
bomad-zamad-kurok-hilab-bimuk-firar-vihav-sisit-filij-busuj-farah-dahif-tibiv-kanad-doviv-gilov-sogot
```
```
$ proquint -base58 bomad-zamad-kurok-hilab-bimuk-firar-vihav-sisit-filij-busuj-farah-dahif-tibiv-kanad-doviv-gilov-sogot
QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB
```

### Base58 and crc32
```
$ proquint -crc32 -base58 QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB
gudis-sisaj-bomad-zamad-kurok-hilab-bimuk-firar-vihav-sisit-filij-busuj-farah-dahif-tibiv-kanad-doviv-gilov-sogot
```

```
$ proquint -crc32 -base58 gudis-sisaj-bomad-zamad-kurok-hilab-bimuk-firar-vihav-sisit-filij-busuj-farah-dahif-tibiv-kanad-doviv-gilov-sogot
QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB
```
