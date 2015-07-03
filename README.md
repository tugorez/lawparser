# lawparser
lawparser se encarga de darle estructura a las leyes y reglamentos ( no sé si funciona igual con otro tipo de documentos legales).

Las leyes y reglamentos tienen una estructura ( no es un estandar escrito para ser seguido por imposición pero casi todos los documentos lo siguen (mas bien un estandar por convención)). La estructura es la siguiente:

	*Todo documento legal puede tener libros
	*Todo libro puede tener títulos
	*Todo titulo puede tener capítulos
	*Todo capítulo puede tener secciones
	*Toda sección tiene artículos
	*Todo artículo puede tener fracciones
	*Toda fracción puede tener subfracciones
(nada facil eh!)
	
Bajo las premisas antes citadas, se elabora la salida del documento legal (descrita en law.go)
	
La documentación del lexer (diagrama de estados) y el parser( cual es la estructura del documento de salida) se encuentran en la carpeta doc.
#sobre los documentos
Regularmente uso calibre (ebook-convert doc.pdf doc.txt) y para docs 
#Uso
en example hay un caso de uso 

