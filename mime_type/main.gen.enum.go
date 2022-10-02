//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package mime_type

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/slices"
)

var (
	All = []MimeType{
		ApplicationAndrewInset,
		ApplicationApplixware,
		ApplicationAtomXml,
		ApplicationAtomcatXml,
		ApplicationAtomsvcXml,
		ApplicationCcxmlXml,
		ApplicationCdmiCapability,
		ApplicationCdmiContainer,
		ApplicationCdmiDomain,
		ApplicationCdmiObject,
		ApplicationCdmiQueue,
		ApplicationCuSeeme,
		ApplicationDavmountXml,
		ApplicationDsscDer,
		ApplicationDsscXml,
		ApplicationEcmascript,
		ApplicationEmmaXml,
		ApplicationEpubZip,
		ApplicationExi,
		ApplicationFontTdpfr,
		ApplicationGpxXml,
		ApplicationGzip,
		ApplicationHyperstudio,
		ApplicationIpfix,
		ApplicationJavaArchive,
		ApplicationJavaSerializedObject,
		ApplicationJavaVm,
		ApplicationJavascript,
		TextJavascript,
		ApplicationJson,
		ApplicationLdJson,
		ApplicationMacBinhex40,
		ApplicationMacCompactpro,
		ApplicationMadsXml,
		ApplicationMarc,
		ApplicationMarcxmlXml,
		ApplicationMathematica,
		ApplicationMathmlXml,
		ApplicationMbox,
		ApplicationMediaservercontrolXml,
		ApplicationMetalink4Xml,
		ApplicationMetsXml,
		ApplicationModsXml,
		ApplicationMp21,
		ApplicationMsword,
		ApplicationMxf,
		ApplicationOctetStream,
		ApplicationOda,
		ApplicationOebpsPackageXml,
		ApplicationOgg,
		ApplicationOnenote,
		ApplicationPatchOpsErrorXml,
		ApplicationPdf,
		ApplicationPgpEncrypted,
		ApplicationPgpSignature,
		ApplicationPicsRules,
		ApplicationPkcs10,
		ApplicationPkcs7Mime,
		ApplicationPkcs7Signature,
		ApplicationPkcs8,
		ApplicationPkixAttrCert,
		ApplicationPkixCert,
		ApplicationPkixCrl,
		ApplicationPkixPkipath,
		ApplicationPkixcmp,
		ApplicationPlsXml,
		ApplicationPostscript,
		ApplicationPrsCww,
		ApplicationPskcXml,
		ApplicationRdfXml,
		ApplicationReginfoXml,
		ApplicationRelaxNgCompactSyntax,
		ApplicationResourceListsXml,
		ApplicationResourceListsDiffXml,
		ApplicationRlsServicesXml,
		ApplicationRsdXml,
		ApplicationRssXml,
		ApplicationRtf,
		ApplicationSbmlXml,
		ApplicationScvpCvRequest,
		ApplicationScvpCvResponse,
		ApplicationScvpVpRequest,
		ApplicationScvpVpResponse,
		ApplicationSdp,
		ApplicationSetPaymentInitiation,
		ApplicationSetRegistrationInitiation,
		ApplicationShfXml,
		ApplicationSmilXml,
		ApplicationSparqlQuery,
		ApplicationSparqlResultsXml,
		ApplicationSrgs,
		ApplicationSrgsXml,
		ApplicationSruXml,
		ApplicationSsmlXml,
		ApplicationTeiXml,
		ApplicationThraudXml,
		ApplicationTimestampedData,
		ApplicationVnd3gppPicBwLarge,
		ApplicationVnd3gppPicBwSmall,
		ApplicationVnd3gppPicBwVar,
		ApplicationVnd3gpp2Tcap,
		ApplicationVnd3mPostItNotes,
		ApplicationVndAccpacSimplyAso,
		ApplicationVndAccpacSimplyImp,
		ApplicationVndAcucobol,
		ApplicationVndAcucorp,
		ApplicationVndAdobeAirApplicationInstallerPackageZip,
		ApplicationVndAdobeFxp,
		ApplicationVndAdobeXdpXml,
		ApplicationVndAdobeXfdf,
		ApplicationVndAheadSpace,
		ApplicationVndAirzipFilesecureAzf,
		ApplicationVndAirzipFilesecureAzs,
		ApplicationVndAmazonEbook,
		ApplicationVndAmericandynamicsAcc,
		ApplicationVndAmigaAmi,
		ApplicationVndAndroidPackageArchive,
		ApplicationVndAnserWebCertificateIssueInitiation,
		ApplicationVndAnserWebFundsTransferInitiation,
		ApplicationVndAntixGameComponent,
		ApplicationVndAppleInstallerXml,
		ApplicationVndAppleMpegurl,
		ApplicationVndAristanetworksSwi,
		ApplicationVndAudiograph,
		ApplicationVndBlueiceMultipass,
		ApplicationVndBmi,
		ApplicationVndBusinessobjects,
		ApplicationVndChemdrawXml,
		ApplicationVndChipnutsKaraokeMmd,
		ApplicationVndCinderella,
		ApplicationVndClaymore,
		ApplicationVndCloantoRp9,
		ApplicationVndClonkC4group,
		ApplicationVndCluetrustCartomobileConfig,
		ApplicationVndCluetrustCartomobileConfigPkg,
		ApplicationVndCommonspace,
		ApplicationVndContactCmsg,
		ApplicationVndCosmocaller,
		ApplicationVndCrickClicker,
		ApplicationVndCrickClickerKeyboard,
		ApplicationVndCrickClickerPalette,
		ApplicationVndCrickClickerTemplate,
		ApplicationVndCrickClickerWordbank,
		ApplicationVndCriticaltoolsWbsXml,
		ApplicationVndCtcPosml,
		ApplicationVndCupsPpd,
		ApplicationVndCurlCar,
		ApplicationVndCurlPcurl,
		ApplicationVndDataVisionRdz,
		ApplicationVndDenovoFcselayoutLink,
		ApplicationVndDna,
		ApplicationVndDolbyMlp,
		ApplicationVndDpgraph,
		ApplicationVndDreamfactory,
		ApplicationVndDvbAit,
		ApplicationVndDvbService,
		ApplicationVndDynageo,
		ApplicationVndEcowinChart,
		ApplicationVndEnliven,
		ApplicationVndEpsonEsf,
		ApplicationVndEpsonMsf,
		ApplicationVndEpsonQuickanime,
		ApplicationVndEpsonSalt,
		ApplicationVndEpsonSsf,
		ApplicationVndEszigno3Xml,
		ApplicationVndEzpixAlbum,
		ApplicationVndEzpixPackage,
		ApplicationVndFdf,
		ApplicationVndFdsnSeed,
		ApplicationVndFlographit,
		ApplicationVndFluxtimeClip,
		ApplicationVndFramemaker,
		ApplicationVndFrogansFnc,
		ApplicationVndFrogansLtf,
		ApplicationVndFscWeblaunch,
		ApplicationVndFujitsuOasys,
		ApplicationVndFujitsuOasys2,
		ApplicationVndFujitsuOasys3,
		ApplicationVndFujitsuOasysgp,
		ApplicationVndFujitsuOasysprs,
		ApplicationVndFujixeroxDdd,
		ApplicationVndFujixeroxDocuworks,
		ApplicationVndFujixeroxDocuworksBinder,
		ApplicationVndFuzzysheet,
		ApplicationVndGenomatixTuxedo,
		ApplicationVndGeogebraFile,
		ApplicationVndGeogebraTool,
		ApplicationVndGeometryExplorer,
		ApplicationVndGeonext,
		ApplicationVndGeoplan,
		ApplicationVndGeospace,
		ApplicationVndGmx,
		ApplicationVndGoogleEarthKmlXml,
		ApplicationVndGoogleEarthKmz,
		ApplicationVndGrafeq,
		ApplicationVndGrooveAccount,
		ApplicationVndGrooveHelp,
		ApplicationVndGrooveIdentityMessage,
		ApplicationVndGrooveInjector,
		ApplicationVndGrooveToolMessage,
		ApplicationVndGrooveToolTemplate,
		ApplicationVndGrooveVcard,
		ApplicationVndHalXml,
		ApplicationVndHandheldEntertainmentXml,
		ApplicationVndHbci,
		ApplicationVndHheLessonPlayer,
		ApplicationVndHpHpgl,
		ApplicationVndHpHpid,
		ApplicationVndHpHps,
		ApplicationVndHpJlyt,
		ApplicationVndHpPcl,
		ApplicationVndHpPclxl,
		ApplicationVndHydrostatixSofData,
		ApplicationVndHzn3dCrossword,
		ApplicationVndIbmMinipay,
		ApplicationVndIbmModcap,
		ApplicationVndIbmRightsManagement,
		ApplicationVndIbmSecureContainer,
		ApplicationVndIccprofile,
		ApplicationVndIgloader,
		ApplicationVndImmervisionIvp,
		ApplicationVndImmervisionIvu,
		ApplicationVndInsorsIgm,
		ApplicationVndInterconFormnet,
		ApplicationVndIntergeo,
		ApplicationVndIntuQbo,
		ApplicationVndIntuQfx,
		ApplicationVndIpunpluggedRcprofile,
		ApplicationVndIrepositoryPackageXml,
		ApplicationVndIsXpr,
		ApplicationVndIsacFcs,
		ApplicationVndJam,
		ApplicationVndJcpJavameMidletRms,
		ApplicationVndJisp,
		ApplicationVndJoostJodaArchive,
		ApplicationVndKahootz,
		ApplicationVndKdeKarbon,
		ApplicationVndKdeKchart,
		ApplicationVndKdeKformula,
		ApplicationVndKdeKivio,
		ApplicationVndKdeKontour,
		ApplicationVndKdeKpresenter,
		ApplicationVndKdeKspread,
		ApplicationVndKdeKword,
		ApplicationVndKenameaapp,
		ApplicationVndKidspiration,
		ApplicationVndKinar,
		ApplicationVndKoan,
		ApplicationVndKodakDescriptor,
		ApplicationVndLasLasXml,
		ApplicationVndLlamagraphicsLifeBalanceDesktop,
		ApplicationVndLlamagraphicsLifeBalanceExchangeXml,
		ApplicationVndLotus123,
		ApplicationVndLotusApproach,
		ApplicationVndLotusFreelance,
		ApplicationVndLotusNotes,
		ApplicationVndLotusOrganizer,
		ApplicationVndLotusScreencam,
		ApplicationVndLotusWordpro,
		ApplicationVndMacportsPortpkg,
		ApplicationVndMcd,
		ApplicationVndMedcalcdata,
		ApplicationVndMediastationCdkey,
		ApplicationVndMfer,
		ApplicationVndMfmp,
		ApplicationVndMicrografxFlo,
		ApplicationVndMicrografxIgx,
		ApplicationVndMif,
		ApplicationVndMobiusDaf,
		ApplicationVndMobiusDis,
		ApplicationVndMobiusMbk,
		ApplicationVndMobiusMqy,
		ApplicationVndMobiusMsl,
		ApplicationVndMobiusPlc,
		ApplicationVndMobiusTxf,
		ApplicationVndMophunApplication,
		ApplicationVndMophunCertificate,
		ApplicationVndMozillaXulXml,
		ApplicationVndMsArtgalry,
		ApplicationVndMsCabCompressed,
		ApplicationVndMsExcel,
		ApplicationVndMsExcelAddinMacroenabled12,
		ApplicationVndMsExcelSheetBinaryMacroenabled12,
		ApplicationVndMsExcelSheetMacroenabled12,
		ApplicationVndMsExcelTemplateMacroenabled12,
		ApplicationVndMsFontobject,
		ApplicationVndMsHtmlhelp,
		ApplicationVndMsIms,
		ApplicationVndMsLrm,
		ApplicationVndMsOfficetheme,
		ApplicationVndMsPkiSeccat,
		ApplicationVndMsPkiStl,
		ApplicationVndMsPowerpoint,
		ApplicationVndMsPowerpointAddinMacroenabled12,
		ApplicationVndMsPowerpointPresentationMacroenabled12,
		ApplicationVndMsPowerpointSlideMacroenabled12,
		ApplicationVndMsPowerpointSlideshowMacroenabled12,
		ApplicationVndMsPowerpointTemplateMacroenabled12,
		ApplicationVndMsProject,
		ApplicationVndMsWordDocumentMacroenabled12,
		ApplicationVndMsWordTemplateMacroenabled12,
		ApplicationVndMsWorks,
		ApplicationVndMsWpl,
		ApplicationVndMsXpsdocument,
		ApplicationVndMseq,
		ApplicationVndMusician,
		ApplicationVndMuveeStyle,
		ApplicationVndNeurolanguageNlu,
		ApplicationVndNoblenetDirectory,
		ApplicationVndNoblenetSealer,
		ApplicationVndNoblenetWeb,
		ApplicationVndNokiaNGageData,
		ApplicationVndNokiaNGageSymbianInstall,
		ApplicationVndNokiaRadioPreset,
		ApplicationVndNokiaRadioPresets,
		ApplicationVndNovadigmEdm,
		ApplicationVndNovadigmEdx,
		ApplicationVndNovadigmExt,
		ApplicationVndOasisOpendocumentChart,
		ApplicationVndOasisOpendocumentChartTemplate,
		ApplicationVndOasisOpendocumentDatabase,
		ApplicationVndOasisOpendocumentFormula,
		ApplicationVndOasisOpendocumentFormulaTemplate,
		ApplicationVndOasisOpendocumentGraphics,
		ApplicationVndOasisOpendocumentGraphicsTemplate,
		ApplicationVndOasisOpendocumentImage,
		ApplicationVndOasisOpendocumentImageTemplate,
		ApplicationVndOasisOpendocumentPresentation,
		ApplicationVndOasisOpendocumentPresentationTemplate,
		ApplicationVndOasisOpendocumentSpreadsheet,
		ApplicationVndOasisOpendocumentSpreadsheetTemplate,
		ApplicationVndOasisOpendocumentText,
		ApplicationVndOasisOpendocumentTextMaster,
		ApplicationVndOasisOpendocumentTextTemplate,
		ApplicationVndOasisOpendocumentTextWeb,
		ApplicationVndOlpcSugar,
		ApplicationVndOmaDd2Xml,
		ApplicationVndOpenofficeorgExtension,
		ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation,
		ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlide,
		ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideshow,
		ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTemplate,
		ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet,
		ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTemplate,
		ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument,
		ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlTemplate,
		ApplicationVndOsgeoMapguidePackage,
		ApplicationVndOsgiDp,
		ApplicationVndPalm,
		ApplicationVndPawaafile,
		ApplicationVndPgFormat,
		ApplicationVndPgOsasli,
		ApplicationVndPicsel,
		ApplicationVndPmiWidget,
		ApplicationVndPocketlearn,
		ApplicationVndPowerbuilder6,
		ApplicationVndPreviewsystemsBox,
		ApplicationVndProteusMagazine,
		ApplicationVndPublishareDeltaTree,
		ApplicationVndPviPtid1,
		ApplicationVndQuarkQuarkxpress,
		ApplicationVndRar,
		ApplicationXRarCompressed,
		ApplicationVndRealvncBed,
		ApplicationVndRecordareMusicxml,
		ApplicationVndRecordareMusicxmlXml,
		ApplicationVndRigCryptonote,
		ApplicationVndRimCod,
		ApplicationVndRnRealmedia,
		ApplicationVndRoute66Link66Xml,
		ApplicationVndSailingtrackerTrack,
		ApplicationVndSeemail,
		ApplicationVndSema,
		ApplicationVndSemd,
		ApplicationVndSemf,
		ApplicationVndShanaInformedFormdata,
		ApplicationVndShanaInformedFormtemplate,
		ApplicationVndShanaInformedInterchange,
		ApplicationVndShanaInformedPackage,
		ApplicationVndSimtechMindmapper,
		ApplicationVndSmaf,
		ApplicationVndSmartTeacher,
		ApplicationVndSolentSdkmXml,
		ApplicationVndSpotfireDxp,
		ApplicationVndSpotfireSfs,
		ApplicationVndStardivisionCalc,
		ApplicationVndStardivisionDraw,
		ApplicationVndStardivisionImpress,
		ApplicationVndStardivisionMath,
		ApplicationVndStardivisionWriter,
		ApplicationVndStardivisionWriterGlobal,
		ApplicationVndStepmaniaStepchart,
		ApplicationVndSunXmlCalc,
		ApplicationVndSunXmlCalcTemplate,
		ApplicationVndSunXmlDraw,
		ApplicationVndSunXmlDrawTemplate,
		ApplicationVndSunXmlImpress,
		ApplicationVndSunXmlImpressTemplate,
		ApplicationVndSunXmlMath,
		ApplicationVndSunXmlWriter,
		ApplicationVndSunXmlWriterGlobal,
		ApplicationVndSunXmlWriterTemplate,
		ApplicationVndSusCalendar,
		ApplicationVndSvd,
		ApplicationVndSymbianInstall,
		ApplicationVndSyncmlXml,
		ApplicationVndSyncmlDmWbxml,
		ApplicationVndSyncmlDmXml,
		ApplicationVndTaoIntentModuleArchive,
		ApplicationVndTmobileLivetv,
		ApplicationVndTridTpt,
		ApplicationVndTriscapeMxs,
		ApplicationVndTrueapp,
		ApplicationVndUfdl,
		ApplicationVndUiqTheme,
		ApplicationVndUmajin,
		ApplicationVndUnity,
		ApplicationVndUomlXml,
		ApplicationVndVcx,
		ApplicationVndVisio,
		ApplicationVndVisio2013,
		ApplicationVndVisionary,
		ApplicationVndVsf,
		ApplicationVndWapWbxml,
		ApplicationVndWapWmlc,
		ApplicationVndWapWmlscriptc,
		ApplicationVndWebturbo,
		ApplicationVndWolframPlayer,
		ApplicationVndWordperfect,
		ApplicationVndWqd,
		ApplicationVndWtStf,
		ApplicationVndXara,
		ApplicationVndXfdl,
		ApplicationVndYamahaHvDic,
		ApplicationVndYamahaHvScript,
		ApplicationVndYamahaHvVoice,
		ApplicationVndYamahaOpenscoreformat,
		ApplicationVndYamahaOpenscoreformatOsfpvgXml,
		ApplicationVndYamahaSmafAudio,
		ApplicationVndYamahaSmafPhrase,
		ApplicationVndYellowriverCustomMenu,
		ApplicationVndZul,
		ApplicationVndZzazzDeckXml,
		ApplicationVoicexmlXml,
		ApplicationWidget,
		ApplicationWinhlp,
		ApplicationWsdlXml,
		ApplicationWspolicyXml,
		ApplicationX7zCompressed,
		ApplicationXAbiword,
		ApplicationXAceCompressed,
		ApplicationXAppleDiskimage,
		ApplicationXAuthorwareBin,
		ApplicationXAuthorwareMap,
		ApplicationXAuthorwareSeg,
		ApplicationXBcpio,
		ApplicationXBittorrent,
		ApplicationXBzip,
		ApplicationXBzip2,
		ApplicationXCdf,
		ApplicationXCdlink,
		ApplicationXChat,
		ApplicationXChessPgn,
		ApplicationXCpio,
		ApplicationXCsh,
		ApplicationXDebianPackage,
		ApplicationXDirector,
		ApplicationXDoom,
		ApplicationXDtbncxXml,
		ApplicationXDtbookXml,
		ApplicationXDtbresourceXml,
		ApplicationXDvi,
		ApplicationXFontBdf,
		ApplicationXFontGhostscript,
		ApplicationXFontLinuxPsf,
		ApplicationXFontPcf,
		ApplicationXFontSnf,
		ApplicationXFontType1,
		ApplicationXFreearc,
		ApplicationXFuturesplash,
		ApplicationXGnumeric,
		ApplicationXGtar,
		ApplicationXHdf,
		ApplicationXHttpdPhp,
		ApplicationXJavaJnlpFile,
		ApplicationXLatex,
		ApplicationXMobipocketEbook,
		ApplicationXMsApplication,
		ApplicationXMsWmd,
		ApplicationXMsWmz,
		ApplicationXMsXbap,
		ApplicationXMsaccess,
		ApplicationXMsbinder,
		ApplicationXMscardfile,
		ApplicationXMsclip,
		ApplicationXMsdownload,
		ApplicationXMsmediaview,
		ApplicationXMsmetafile,
		ApplicationXMsmoney,
		ApplicationXMspublisher,
		ApplicationXMsschedule,
		ApplicationXMsterminal,
		ApplicationXMswrite,
		ApplicationXNetcdf,
		ApplicationXPkcs12,
		ApplicationXPkcs7Certificates,
		ApplicationXPkcs7Certreqresp,
		ApplicationXSh,
		ApplicationXShar,
		ApplicationXShockwaveFlash,
		ApplicationXSilverlightApp,
		ApplicationXStuffit,
		ApplicationXStuffitx,
		ApplicationXSv4cpio,
		ApplicationXSv4crc,
		ApplicationXTar,
		ApplicationXTcl,
		ApplicationXTex,
		ApplicationXTexTfm,
		ApplicationXTexinfo,
		ApplicationXUstar,
		ApplicationXWaisSource,
		ApplicationXX509CaCert,
		ApplicationXXfig,
		ApplicationXXpinstall,
		ApplicationXYaml,
		TextXYaml,
		TextYaml,
		ApplicationYaml,
		ApplicationXcapDiffXml,
		ApplicationXencXml,
		ApplicationXhtmlXml,
		ApplicationXml,
		TextXml,
		ApplicationXmlDtd,
		ApplicationXopXml,
		ApplicationXsltXml,
		ApplicationXspfXml,
		ApplicationXvXml,
		ApplicationYang,
		ApplicationYinXml,
		ApplicationZip,
		AudioAdpcm,
		AudioBasic,
		AudioMidi,
		AudioXMidi,
		AudioMp4,
		AudioMpeg,
		AudioOgg,
		AudioOpus,
		AudioVndDeceAudio,
		AudioVndDigitalWinds,
		AudioVndDra,
		AudioVndDts,
		AudioVndDtsHd,
		AudioVndLucentVoice,
		AudioVndMsPlayreadyMediaPya,
		AudioVndNueraEcelp4800,
		AudioVndNueraEcelp7470,
		AudioVndNueraEcelp9600,
		AudioVndRip,
		AudioWav,
		AudioXWav,
		AudioWebm,
		AudioXAac,
		AudioXAiff,
		AudioXMpegurl,
		AudioXMsWax,
		AudioXMsWma,
		AudioXPnRealaudio,
		AudioXPnRealaudioPlugin,
		ChemicalXCdx,
		ChemicalXCif,
		ChemicalXCmdf,
		ChemicalXCml,
		ChemicalXCsml,
		ChemicalXXyz,
		FontOtf,
		ApplicationXFontOtf,
		FontTtf,
		ApplicationXFontTtf,
		FontWoff,
		ApplicationXFontWoff,
		FontWoff2,
		ImageBmp,
		ImageCgm,
		ImageG3fax,
		ImageGif,
		ImageIef,
		ImageJpeg,
		ImageXCitrixJpeg,
		ImageKtx,
		ImagePjpeg,
		ImagePng,
		ImageXCitrixPng,
		ImageXPng,
		ImagePrsBtif,
		ImageSvgXml,
		ImageTiff,
		ImageVndAdobePhotoshop,
		ImageVndDeceGraphic,
		ImageVndDjvu,
		ImageVndDvbSubtitle,
		ImageVndDwg,
		ImageVndDxf,
		ImageVndFastbidsheet,
		ImageVndFpx,
		ImageVndFst,
		ImageVndFujixeroxEdmicsMmr,
		ImageVndFujixeroxEdmicsRlc,
		ImageVndMsModi,
		ImageVndNetFpx,
		ImageVndWapWbmp,
		ImageVndXiff,
		ImageWebp,
		ImageXCmuRaster,
		ImageXCmx,
		ImageXFreehand,
		ImageXIcon,
		ImageVndMicrosoftIcon,
		ImageXPcx,
		ImageXPict,
		ImageXPortableAnymap,
		ImageXPortableBitmap,
		ImageXPortableGraymap,
		ImageXPortablePixmap,
		ImageXRgb,
		ImageXXbitmap,
		ImageXXpixmap,
		ImageXXwindowdump,
		MessageRfc822,
		ModelIges,
		ModelMesh,
		ModelVndColladaXml,
		ModelVndDwf,
		ModelVndGdl,
		ModelVndGtw,
		ModelVndMts,
		ModelVndVtu,
		ModelVrml,
		TextCalendar,
		TextCss,
		TextCsv,
		TextHtml,
		TextN3,
		TextPlain,
		TextPlainBas,
		TextPrsLinesTag,
		TextRichtext,
		TextSgml,
		TextTabSeparatedValues,
		TextTroff,
		TextTurtle,
		TextUriList,
		TextVndCurl,
		TextVndCurlDcurl,
		TextVndCurlMcurl,
		TextVndCurlScurl,
		TextVndFly,
		TextVndFmiFlexstor,
		TextVndGraphviz,
		TextVndIn3d3dml,
		TextVndIn3dSpot,
		TextVndSunJ2meAppDescriptor,
		TextVndWapWml,
		TextVndWapWmlscript,
		TextXAsm,
		TextXC,
		TextXFortran,
		TextXJavaSource,
		Java,
		TextXPascal,
		TextXSetext,
		TextXUuencode,
		TextXVcalendar,
		TextXVcard,
		Unkown,
		Video3gpp,
		Audio3gpp,
		Video3gpp2,
		Audio3gpp2,
		VideoH261,
		VideoH263,
		VideoH264,
		VideoJpeg,
		VideoJpm,
		VideoMj2,
		VideoMp2t,
		VideoMp4,
		ApplicationMp4,
		VideoMpeg,
		VideoOgg,
		VideoQuicktime,
		VideoVndDeceHd,
		VideoVndDeceMobile,
		VideoVndDecePd,
		VideoVndDeceSd,
		VideoVndDeceVideo,
		VideoVndFvt,
		VideoVndMpegurl,
		VideoVndMsPlayreadyMediaPyv,
		VideoVndUvvuMp4,
		VideoVndVivo,
		VideoWebm,
		VideoXF4v,
		VideoXFli,
		VideoXFlv,
		VideoXM4v,
		VideoXMsAsf,
		VideoXMsWm,
		VideoXMsWmv,
		VideoXMsWmx,
		VideoXMsWvx,
		VideoXMsvideo,
		VideoXSgiMovie,
		XConferenceXCooltalk,
	}
)

func (t MimeType) String() string {
	return string(t)
}

func Parse(v string) (MimeType, error) {
	f, ok := slices.FindFn(All, func(x MimeType) bool {
		return MimeType(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slices.ContainsFn(All, func(v MimeType) bool {
		return string(v) == s
	})
}

var ErrInvalid = errors.New("invalid enumeration type")

func ErrorV(v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrInvalid, v, slices.Join(All, ","),
	)
}

func (t MimeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *MimeType) UnmarshalJSON(data []byte) error {
	var v string

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t MimeType) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *MimeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	if err := unmarshal(&v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
