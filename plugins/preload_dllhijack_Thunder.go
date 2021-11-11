package plugins

// PreloadDllHijackThunder 位于 迅雷11 安装目录的 Thunder.exe
type PreloadDllHijackThunder struct {
	SubPluginPreloadDllHijackX86
}

func (p *PreloadDllHijackThunder) GetMainProgramName() (string) {
	return "Thunder"
}

func (p *PreloadDllHijackThunder) GetDllName() (string) {
	return "libexpat"
}

func (p *PreloadDllHijackThunder) GetPluginName() (string) {
	return "preload_dll_hijack_Thunder"
}

func (p *PreloadDllHijackThunder) GetDllExports() []string {
	return []string{
		"Function",
		"XML_DefaultCurrent",
		"XML_ErrorString",
		"XML_ExpatVersion",
		"XML_ExpatVersionInfo",
		"XML_ExternalEntityParserCreate",
		"XML_GetBase",
		"XML_GetBuffer",
		"XML_GetCurrentByteCount",
		"XML_GetCurrentByteIndex",
		"XML_GetCurrentColumnNumber",
		"XML_GetCurrentLineNumber",
		"XML_GetErrorCode",
		"XML_GetIdAttributeIndex",
		"XML_GetInputContext",
		"XML_GetSpecifiedAttributeCount",
		"XML_Parse",
		"XML_ParseBuffer",
		"XML_ParserCreate",
		"XML_ParserCreateNS",
		"XML_ParserCreate_MM",
		"XML_ParserFree",
		"XML_SetAttlistDeclHandler",
		"XML_SetBase",
		"XML_SetCdataSectionHandler",
		"XML_SetCharacterDataHandler",
		"XML_SetCommentHandler",
		"XML_SetDefaultHandler",
		"XML_SetDefaultHandlerExpand",
		"XML_SetDoctypeDeclHandler",
		"XML_SetElementDeclHandler",
		"XML_SetElementHandler",
		"XML_SetEncoding",
		"XML_SetEndCdataSectionHandler",
		"XML_SetEndDoctypeDeclHandler",
		"XML_SetEndElementHandler",
		"XML_SetEndNamespaceDeclHandler",
		"XML_SetEntityDeclHandler",
		"XML_SetExternalEntityRefHandler",
		"XML_SetExternalEntityRefHandlerArg",
		"XML_SetNamespaceDeclHandler",
		"XML_SetNotStandaloneHandler",
		"XML_SetNotationDeclHandler",
		"XML_SetParamEntityParsing",
		"XML_SetProcessingInstructionHandler",
		"XML_SetReturnNSTriplet",
		"XML_SetStartCdataSectionHandler",
		"XML_SetStartDoctypeDeclHandler",
		"XML_SetStartElementHandler",
		"XML_SetStartNamespaceDeclHandler",
		"XML_SetUnknownEncodingHandler",
		"XML_SetUnparsedEntityDeclHandler",
		"XML_SetUserData",
		"XML_SetXmlDeclHandler",
		"XML_UseParserAsHandlerArg",
		"XML_ParserReset",
		"XML_SetSkippedEntityHandler",
		"XML_GetFeatureList",
		"XML_UseForeignDTD",
		"XML_FreeContentModel",
		"XML_MemMalloc",
		"XML_MemRealloc",
		"XML_MemFree",
		"XML_StopParser",
		"XML_ResumeParser",
		"XML_GetParsingStatus",
	}
}
