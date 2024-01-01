package repository

import "gorm.io/gorm"

func migrateValues(conn *gorm.DB) {
	// dir, err := config.FindModuleDir("domain/sqldomain")
	// if err != nil {
	// 	panic(err)
	// }

	// config := &packages.Config{
	// 	Mode: packages.NeedTypes | packages.NeedSyntax,
	// }
	// pkgs, err := packages.Load(config, dir)

	// for _, syntax := range pkgs[0].Syntax {
	// 	for _, decl := range syntax.Decls {
	// 		if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "GetId" {
	// 			if len(funcDecl.Recv.List) > 0 {
	// 				receiverType := funcDecl.Recv.List[0].Type
	//
	// 				if ident, ok := receiverType.(*ast.Ident); ok {
	//
	// 					// fazer uma instancia e automigrate do gorm
	//
	// 				}
	//
	// 			}
	// 		}
	//
	// 	}
	// }
}
