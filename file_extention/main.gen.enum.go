//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package file_extention

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/slicer" // v1.0.15
)

var (
	All = []FileExtention{
		Aw,
		Atom,
		Atomcat,
		Atomsvc,
		Ccxml,
		Cdmia,
		Cdmic,
		Cdmid,
		Cdmio,
		Cdmiq,
		Cu,
		Davmount,
		Dssc,
		Xdssc,
		Es,
		Emma,
		Epub,
		Exi,
		Pfr,
		Gpx,
		Gz,
		Stk,
		Ipfix,
		Jar,
		Ser,
		Class,
		Js,
		Json,
		Jsonld,
		Hqx,
		Cpt,
		Mads,
		Mrc,
		Mrcx,
		Ma,
		Mathml,
		Mbox,
		Mscml,
		Meta4,
		Mets,
		Mods,
		M21,
		Doc,
		Mxf,
		Bin,
		Oda,
		Opf,
		Ogx,
		Onetoc,
		Xer,
		Pdf,
		Pgp,
		Prf,
		P10,
		P7m,
		P7s,
		P8,
		Ac,
		Cer,
		Crl,
		Pkipath,
		Pki,
		Pls,
		Ai,
		Cww,
		Pskcxml,
		Rdf,
		Rif,
		Rnc,
		Rl,
		Rld,
		Rs,
		Rsd,
		Rss,
		Rtf,
		Sbml,
		Scq,
		Scs,
		Spq,
		Spp,
		Sdp,
		Setpay,
		Setreg,
		Shf,
		Smi,
		Rq,
		Srx,
		Gram,
		Grxml,
		Sru,
		Ssml,
		Tei,
		Tfi,
		Tsd,
		Plb,
		Psb,
		Pvb,
		Tcap,
		Pwn,
		Aso,
		Imp,
		Acu,
		Atc,
		Air,
		Fxp,
		Xdp,
		Xfdf,
		Ahead,
		Azf,
		Azs,
		Azw,
		Acc,
		Ami,
		Apk,
		Cii,
		Fti,
		Atx,
		Mpkg,
		M3u8,
		Swi,
		Aep,
		Mpm,
		Bmi,
		Rep,
		Cdxml,
		Mmd,
		Cdy,
		Cla,
		Rp9,
		C4g,
		C11amc,
		C11amz,
		Csp,
		Cdbcmsg,
		Cmc,
		Clkx,
		Clkk,
		Clkp,
		Clkt,
		Clkw,
		Wbs,
		Pml,
		Ppd,
		Car,
		Pcurl,
		Rdz,
		Fe_launch,
		Dna,
		Mlp,
		Dpg,
		Dfac,
		Ait,
		Svc,
		Geo,
		Mag,
		Nml,
		Esf,
		Msf,
		Qam,
		Slt,
		Ssf,
		Es3,
		Ez2,
		Ez3,
		Fdf,
		Seed,
		Gph,
		Ftc,
		Fm,
		Fnc,
		Ltf,
		Fsc,
		Oas,
		Oa2,
		Oa3,
		Fg5,
		Bh2,
		Ddd,
		Xdw,
		Xbd,
		Fzs,
		Txd,
		Ggb,
		Ggt,
		Gex,
		Gxt,
		G2w,
		G3w,
		Gmx,
		Kml,
		Kmz,
		Gqf,
		Gac,
		Ghf,
		Gim,
		Grv,
		Gtm,
		Tpl,
		Vcg,
		Hal,
		Zmm,
		Hbci,
		Les,
		Hpgl,
		Hpid,
		Hps,
		Jlt,
		Pcl,
		Pclxl,
		SfdHdstx,
		X3d,
		Mpy,
		Afp,
		Irm,
		Sc,
		Icc,
		Igl,
		Ivp,
		Ivu,
		Igm,
		Xpw,
		I2g,
		Qbo,
		Qfx,
		Rcprofile,
		Irp,
		Xpr,
		Fcs,
		Jam,
		Rms,
		Jisp,
		Joda,
		Ktz,
		Karbon,
		Chrt,
		Kfo,
		Flw,
		Kon,
		Kpr,
		Ksp,
		Kwd,
		Htke,
		Kia,
		Kne,
		Skp,
		Sse,
		Lasxml,
		Lbd,
		Lbe,
		_123,
		Apr,
		Pre,
		Nsf,
		Org,
		Scm,
		Lwp,
		Portpkg,
		Mcd,
		Mc1,
		Cdkey,
		Mwf,
		Mfm,
		Flo,
		Igx,
		Mif,
		Daf,
		Dis,
		Mbk,
		Mqy,
		Msl,
		Plc,
		Txf,
		Mpn,
		Mpc,
		Xul,
		Cil,
		Cab,
		Xls,
		Xlam,
		Xlsb,
		Xlsm,
		Xltm,
		Eot,
		Chm,
		Ims,
		Lrm,
		Thmx,
		Cat,
		Stl,
		Ppt,
		Ppam,
		Pptm,
		Sldm,
		Ppsm,
		Potm,
		Mpp,
		Docm,
		Dotm,
		Wps,
		Wpl,
		Xps,
		Mseq,
		Mus,
		Msty,
		Nlu,
		Nnd,
		Nns,
		Nnw,
		Ngdat,
		NGage,
		Rpst,
		Rpss,
		Edm,
		Edx,
		Ext,
		Odc,
		Otc,
		Odb,
		Odf,
		Odft,
		Odg,
		Otg,
		Odi,
		Oti,
		Odp,
		Otp,
		Ods,
		Ots,
		Odt,
		Odm,
		Ott,
		Oth,
		Xo,
		Dd2,
		Oxt,
		Pptx,
		Sldx,
		Ppsx,
		Potx,
		Xlsx,
		Xltx,
		Docx,
		Dotx,
		Mgp,
		Dp,
		Pdb,
		Paw,
		Str,
		Ei6,
		Efif,
		Wg,
		Plf,
		Pbd,
		Box,
		Mgz,
		Qps,
		Ptid,
		Qxd,
		Rar,
		Bed,
		Mxl,
		Musicxml,
		Cryptonote,
		Cod,
		Rm,
		Link66,
		St,
		See,
		Sema,
		Semd,
		Semf,
		Ifm,
		Itp,
		Iif,
		Ipk,
		Twd,
		Mmf,
		Teacher,
		Sdkm,
		Dxp,
		Sfs,
		Sdc,
		Sda,
		Sdd,
		Smf,
		Sdw,
		Sgl,
		Sm,
		Sxc,
		Stc,
		Sxd,
		Std,
		Sxi,
		Sti,
		Sxm,
		Sxw,
		Sxg,
		Stw,
		Sus,
		Svd,
		Sis,
		Xsm,
		Bdm,
		Xdm,
		Tao,
		Tmo,
		Tpt,
		Mxs,
		Tra,
		Ufd,
		Utz,
		Umj,
		Unityweb,
		Uoml,
		Vcx,
		Vsd,
		Vsdx,
		Vis,
		Vsf,
		Wbxml,
		Wmlc,
		Wmlsc,
		Wtb,
		Nbp,
		Wpd,
		Wqd,
		Stf,
		Xar,
		Xfdl,
		Hvd,
		Hvs,
		Hvp,
		Osf,
		Osfpvg,
		Saf,
		Spf,
		Cmp,
		Zir,
		Zaz,
		Vxml,
		Wgt,
		Hlp,
		Wsdl,
		Wspolicy,
		_7z,
		Abw,
		Ace,
		Dmg,
		Aab,
		Aam,
		Aas,
		Bcpio,
		Torrent,
		Bz,
		Bz2,
		Cda,
		Vcd,
		Chat,
		Pgn,
		Cpio,
		Csh,
		Deb,
		Dir,
		Wad,
		Ncx,
		Dtb,
		Res,
		Dvi,
		Bdf,
		Gsf,
		Psf,
		Pcf,
		Snf,
		Pfa,
		Arc,
		Spl,
		Gnumeric,
		Gtar,
		Hdf,
		Php,
		Jnlp,
		Latex,
		Prc,
		Application,
		Wmd,
		Wmz,
		Xbap,
		Mdb,
		Obd,
		Crd,
		Clp,
		Exe,
		Mvb,
		Wmf,
		Mny,
		Pub,
		Scd,
		Trm,
		Wri,
		Nc,
		P12,
		P7b,
		P7r,
		Sh,
		Shar,
		Swf,
		Xap,
		Sit,
		Sitx,
		Sv4cpio,
		Sv4crc,
		Tar,
		Tcl,
		Tex,
		Tfm,
		Texinfo,
		Ustar,
		Src,
		Der,
		Fig,
		Xpi,
		Yaml,
		Yml,
		Xdf,
		Xenc,
		Xhtml,
		Xml,
		Dtd,
		Xop,
		Xslt,
		Xspf,
		Mxml,
		Yang,
		Yin,
		Zip,
		Adp,
		Au,
		Mid,
		Midi,
		Mp4a,
		Mpga,
		Oga,
		Opus,
		Uva,
		Eol,
		Dra,
		Dts,
		Dtshd,
		Lvp,
		Pya,
		Ecelp4800,
		Ecelp7470,
		Ecelp9600,
		Rip,
		Wav,
		Weba,
		Aac,
		Aif,
		M3u,
		Wax,
		Wma,
		Ram,
		Rmp,
		Cdx,
		Cif,
		Cmdf,
		Cml,
		Csml,
		Xyz,
		Otf,
		Ttf,
		Woff,
		Woff2,
		Bmp,
		Cgm,
		G3,
		Gif,
		Ief,
		Jpeg,
		Jpg,
		Ktx,
		Pjpeg,
		Png,
		Btif,
		Svg,
		Tif,
		Tiff,
		Psd,
		Uvi,
		Djvu,
		Sub,
		Dwg,
		Dxf,
		Fbs,
		Fpx,
		Fst,
		Mmr,
		Rlc,
		Mdi,
		Npx,
		Wbmp,
		Xif,
		Webp,
		Ras,
		Cmx,
		Fh,
		Ico,
		Pcx,
		Pic,
		Pnm,
		Pbm,
		Pgm,
		Ppm,
		Rgb,
		Xbm,
		Xpm,
		Xwd,
		Eml,
		Igs,
		Msh,
		Dae,
		Dwf,
		Gdl,
		Gtw,
		Mts,
		Vtu,
		Wrl,
		Ics,
		Css,
		Csv,
		Html,
		N3,
		Txt,
		Par,
		Dsc,
		Rtx,
		Sgml,
		Tsv,
		T,
		Ttl,
		Uri,
		Curl,
		Dcurl,
		Mcurl,
		Scurl,
		Fly,
		Flx,
		Gv,
		_3dml,
		Spot,
		Jad,
		Wml,
		Wmls,
		S,
		C,
		F,
		Java,
		P,
		Etx,
		Uu,
		Vcs,
		Vcf,
		_3gp,
		_3g2,
		H261,
		H263,
		H264,
		Jpgv,
		Jpm,
		Mj2,
		Ts,
		Mp4,
		Mpeg,
		Ogv,
		Qt,
		Uvh,
		Uvm,
		Uvp,
		Uvs,
		Uvv,
		Fvt,
		Mxu,
		Pyv,
		Uvu,
		Viv,
		Webm,
		F4v,
		Fli,
		Flv,
		M4v,
		Asf,
		Wm,
		Wmv,
		Wmx,
		Wvx,
		Avi,
		Movie,
		Ice,
		Toml,
		Hcl,
		Ini,
	}
)

func (t FileExtention) String() string {
	return string(t)
}

func Parse(v string) (FileExtention, error) {
	f, ok := slicer.FindFn(All, func(x FileExtention) bool {
		return FileExtention(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slicer.ContainsFn(All, func(v FileExtention) bool {
		return string(v) == s
	})
}

var ErrInvalid = errors.New("invalid enumeration type")

func ErrorV(v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrInvalid, v, slicer.Join(All, ","),
	)
}

func (t FileExtention) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *FileExtention) UnmarshalJSON(data []byte) error {
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

func (t FileExtention) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *FileExtention) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
