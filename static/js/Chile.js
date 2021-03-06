(function (root, factory) {
    if (typeof define === 'function' && define.amd) {
        define(['exports', 'echarts'], factory);
    } else if (typeof exports === 'object' && typeof exports.nodeName !== 'string') {
        factory(exports, require('echarts'));
    } else {
        factory({}, root.echarts);
    }
}(this, function (exports, echarts) {
    var log = function (msg) {
        if (typeof console !== 'undefined') {
            console && console.error && console.error(msg);
        }
    };
    if (!echarts) {
        log('ECharts is not Loaded');
        return;
    }
    if (!echarts.registerMap) {
        log('ECharts Map is not loaded');
        return;
    }
    echarts.registerMap('智利', {
        "type": "FeatureCollection",
        "features": [{
            "type": "Feature",
            "properties": {"name": "III Región de Atacama"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@Î@VfeöntfĒ¤Rx~p\\ze¸cà@ªĬr®©nöJ´q¼DĐO@~vLº´m®jxW´Zrpê^fnÎ_l[î^ºKn~rÄJ^U XðLjÆZf RvWrCkUhy_~gIcreuAnGPĈЅƩĝOkƣ˂΋awl¥§Aĭ¯q§rcKaÃT[ĳAiSaDYw{qµÏ}^kitcjs{}O£`¥{B¿gÁc[EkuoyP³FoEs©sIT{jãiÃćp_£`SCčyaOsTÅdKYå{iC}kZÍGē£DQ·§ey^{y±F«ª]t_DHy¯WďbÊéä·By»bWzrlfnQ~obA´SVrËdóMWZq_i­_aeE¹X¹ĭK¡m`MÉRuXaĊÅPaOÃoAcwQÙH]ĊLtoXXzS~¼f@ò²NXw¶¤DÆZbrV°tAdÔWCÂ_^StZĄ´M~vÒÈGjrnsð¢`jQÌlÈW}¼iAÆe\\b®t|vK~nCczRpYilvj¸te¤|fH¢^bT¨KÆhº]~¶`WnvmAÚbrI´"],
                "encodeOffsets": [[-72354, -26686]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "II Región de Antofagasta"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@uXQeÅYiïKW]VÃI}qmL¹]\\í`kmÍeé]qo³YwX­in¹³uK@}ďP»C³rõIm­ªīq©ß@·dyfo[}w£QđseõmfUeÍ@nRR²g ¡ªa¶sb^xnJÞR^fQ|D°b²AÌ@Æ^hOÐh@¸w¼fnlÊCt·ÌNmvA¾ÀVxOpÆob®IhWZ~hĔB¸dLa¬jdgøF¤fÐgdOd¦z^ðXºSÎpHi­gc¿hÎ_BpepC PYTvcr¨¦P`©¸Lvz~Ēb^@qÔHhbZÐM~pXeÀx\\ªwznðW^xňF¶ZÚrMXb\\ÖSdz~TUŚPºtüQnÒWÎZĢgFxh°UJe¶×Å^UȢȞϤĦ˞Ǿ ĸºD^VİƍĮǣAʩÀ×t»vyHuOƿĄÕ­e÷]K»X­z[PB¿n|EYåmuGóJ½M[Ȣi͎òǆƽÓȍƋыćʧѝǟ΅Ƈi}g[ÛNõK{eHXµ}Ý~uN¤lL¦áJµxYÑÁ}«dUmDeKcÍ"],
                "encodeOffsets": [[-70190, -25905]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Amazonas"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@LPҒźz[¦Frp~@zPNZlQjÑfR²e¨dÀknlÚØãþlWDrt¶ Fh¼WÚ´D|YÖ­ glĐĜdDh§È}ôXvYxXÊhJÈZ²z¤\\hf^ÖMHþIâjv]ÆX¸Ø¦¤´L°dÊL¤tzzf[Iq]èpsT±Zô½ItlX¼§`y¯¸¯\\GZ¾vFHv_ltelÀ t^`dE\\tN~|vEVrxA¥{}{_sUdsVïS^ueH¯u~³ASv]bressFwYhHdC¥EEttJ|\\dU¦|lqyXuwh`D¤x@¯YeÀrN o¨ÇBAÃLOJ¦v|EtmxHbKÚPr@x_cãp¯E}ZF­oLëЃßνčҿíЇǯą҃y`áâkÃ ¢uÄ»vÖÁpuqCÏc®cTcN{yAǰ˪ƖɞɊͺhQDYjuWO}LHSwXSÂc|eFq}IqnddLVx³`Wuq\\Qenm{mKD}Yvb¨¥E\\¢pÃfMp£@uXiºDu}cAEEyzµkEą^LSmO]oCY@[f^ck^cjwW[J ­`}@O{mgRV¬{B__kkhscÅ_cLQ}{Eq|}Cg_ÁqFC±i^{DeWrKŅ]rVEyyi@qgOG}qs±{Jm@ibKÈqme«ujeÄ¢jRrCpVl~¦b¬IxPVJ¶£bay¡Fq©~_Cn{|J~f~H~¬¨wfZ Yl@tÚEl}lQ°{EsvU{Z{mąGvqzQ¯MKd^°kQ½jx«E¡|Vc}FHSDwt\\B«KdfLpp\\pkK"],
                "encodeOffsets": [[-76172, -572]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Antioquia"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@|ċ¦rÓpVO~ęDÅb³{mqđiYqËumU­ZiIïeUģ@»_KX¾ěpeΌŀLj°ArddÈTÚhT~{`a~D\\Ê¼nÎR¢\\ÈV¾jGÀ¸|²QĊL¢IÞgªÄP¼ AZ¯¬_Ðry¼w·u`GÅpó{ÃHoaSÅD¿mOO§\\oÎ«hzBnxDLâ¢Cr«kaG^cEi¿ÝCuhcEÏPf_vRMˤʖUu@µ]nÑtuX¿£­Û¿½}±ÕÉ¥W©iiu@Qd«GzyFk¯_u³¡§qS©½E}­RsMFepÓ¡mQ»YµMoÉÇE³¤ÍY§_ÕeIÝyga½Qi]«å[yTA¤}ÎxTO\\©n¥ÅBZ³\\@¯tSQ\\ēwIk\\ÁO¥Y_smP[v§VsLQ|±ildA²ihM¨zæFwePQmQÊ¾±QéKw«R_]dA¡\\oÛJBÃÚGÀhq^G|f¦iGWBEisPF¸|I´XfzC¬efx~DVzY¸gZIŮngzPppL¾îC sº±µ|©mJíĞwzÕÆbSlµ¼mGIÄ²`GväVNtRvrMxpTXxOFg­yQéTsYòAXxÌWj|M}`aâcf¤EÈWeJe¨Ï~âĞZün^\\dNÒptFhv¨v^"],
                "encodeOffsets": [[-78247, 9104]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Arauca"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@|\\pI^|D^|vKFª`ÖlMpFädZ¢ülÀ~Ģ`pÄD¸DnhsĐ@c~tEZcAÈ`ØYÄZdò]|PLGpôÆDÞ\\ô}bhMÉÖB¸wx¨UZzØRǦϏ˴ΩosçāMïHédwaGuT[¤}fWn@l¢±|cÙfçEõb§SIcS©a^ƅoF±|¿w©PÍKut«Bype«lHgkÇJÍV{·]\\]|gC}«PďKwqGÛV½±YQxÍvabK¶»|oªD"],
                "encodeOffsets": [[-74101, 6459]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Archipiélago de San Andrés, Providencia y Santa Catalina"},
            "geometry": {
                "type": "MultiPolygon",
                "coordinates": [["@@VhCGuL"], ["@@TvvBEV"]],
                "encodeOffsets": [[[-83694, 12839]], [[-83347, 13665]]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Bolívar"},
            "geometry": {
                "type": "MultiPolygon",
                "coordinates": [["@@HÎÏLÉkcJW­ ßMx¡HƁEw¿FÅa_HÙruD\\J£d¥åYïuwaj×[IÑˣʕNuQe`OFÐgdDvÞjÀF]dHlb¬q¡DKáwCmyAgÍ¬[pP¨nPCÀTÆpbGÄ|oôHÆ_v¸»xqzÏ«`Y°ÞìW~xrIz°HlF N¦{hLK²_¾hj[cni°[Fa±LlTehqÆi|J®`nOjB´ÇAbAÖDµÀdkHuUZ`H¢VVT¨YNvÎAÕ\\^]z`_~CQifĒµQAºaPi{B^¤|K¢NrE¼fG¢ir´R¤NhĄfr^JÂRpa@sÑHkY¹Z]y½äaCi±¦Z²tK²ÑėDu|]Kna~voG½agcfµL³Kz¡acnFÁy\\YE×rPxeNqFxUXsYR¨ÉrKrD¹fÀ{EdUPÒó]AVhOXCctHtAo"], ["@@hckgP"]],
                "encodeOffsets": [[[-75642, 9132]], [[-77407, 10596]]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Casanare"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@Cp©¼{LµbaÎu wR²Z¾ÜUrHxĐL¬O~hD{^ [¸^|ÎUÈIhlG¬kfzo¬AvsÎLªOÀx²{EƆpb]ªTJd¨TöaèFÚed²{¡k@Xm~e\\£vSHeic£{oÝęcµ§I}ëĉçyue©U¿Ï»«kÕ¯uwÕGX½[áË¥çk·ĩ¿wmI¥ë_{{{]@ÓiQY·S}şÃXgo{ýęáEe³lµW£HayoYj©xAGyNYbw{¥l{¸{^_wx~ÓĒOjÉÌaFJf}²ÄI~èy ÖpjK¹p®¶ĜnS|zNz«ehd  dNrrVt nL©xXĀfØU¹GbXzR¤cDmzRAtuv@anÌl"],
                "encodeOffsets": [[-74101, 6459]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Cesar"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@¸RÄpð@²Kv_lbFµm}¼SNs¢Jhm²A«U§µwi\\¢zk¦G¶~FR¥iTUV¥qáTs¡D³]`_Z­tc£¯U¥¿e§WÑC£­Jg[\\m§auO¥eIYÁ_Hg±Ma¯V_A«kµZÛidoVÍiE{uGumayXqM¡PÇ½zgpTQPssNsm¼^eÜrN|a\\§K»iWJDimKt}WB©R´Fy@©u{OQV¡Åûs£e~}Nz¥Pį@^zW âÊG¨Lho®qS°EÆxÀFGƂw¢NàX®IldKÊÍÐGÀPvÅ®uÂËXM^nN¢vRÚÉoxÁ t ÈȺªG¢sPp]§ÔmxVa¶`nKrR tLzrVTzvDv^BrX¬KIJ~KzÎ^ScLSú]zfvy¢d~"],
                "encodeOffsets": [[-75368, 11096]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Córdoba"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@Èpx²^¸lXĄ¦rhlInxÈhpjzShnriİxMN}_oÀ}^ffGf®Ac[`o_CwXe¦Gu¤LrqG{aaGy·uJUJÅ[ëR¥jQ}^Õ¸e|²pâ ¶X^rGf®lOahÝBO»©ÃhÝ¡JĉKR{±¿·iHU½Ç[Q¡Í»m[É}Cb_}|gSSÙcÇcBqi¯ĿK΋of½ĜWL`@¼VĤfJðYjV®vnrÌjZrĒn ´|aCÆĚ}PoUqÔ¥{Č"],
                "encodeOffsets": [[-78247, 9104]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Cundinamarca"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@Îf^@raÌYp´MxZpZ[XĂ}´hB`¡dmKjYG}b qR]W©dSf£ÄSxæuZÞY¶ú~IP§²iÎ½ÂcAÏ`wcd·Ebh@Q¡^­aB{Jhi°`¼[¼ñqJF¨WYT]àmLz°FnhEAAÓЍWÉ¸ĭ~bGÛÏvYª]bE¤{Fv|Ohda¥}q{EwWikGsdePß_oT£½Req}K{JWy}_¡CSA¥Of¼E¼lVCw]¢ªú]B FHXpLR¬yndYĪLzÇ`Oq}um¯gE¡§}HÕdU¯Ņ@ċXwoWáY£\\vīYyf_QÓoQ­íRUY@M\\l@wj\\\\ÌkH lK¾eXg¬­kMy³_§XYryDÓgWrPxxdtôêe|LÄcZ~ijHU\\Npt\\Iebúa`fvU\\v@Zf]´b]tLrU¢"],
                "encodeOffsets": [[-76527, 5414]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Huila"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@Fx\\\\ü§xuD¸Lx`¦J¶mpÌWX`GpzxKdZP|¤O¾°ÀM®uVvPfeÆøEJmCAÇNw°UøZ¸ªÄ¢YPPVjSzn¬ÊÎlÀG^xi@k[NZ@nO¡«uéwmD¥·pEwolOu{i­gWYEsw[k§Uoċĕ­çqkÅQX{eZË{]FL{q{q©sGWWÛg[^yásyÉeM]Õ£ÝµO@¡sY¹qwJ£ik]q¤B«Gixg©IM\\b²Wf{qoÊxñly@pIZfVMT|OzÊ]\\ibAÐ~^\\¤zd¶Zz]²»rYFvsÌjò¸KT¼ºNqpe|JHj~DzcvAzÂBhJ¼ėĕðÅÊ"],
                "encodeOffsets": [[-77853, 3012]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "La Guajira"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@Eµ}¥Hl¡y[xj¨¶VB¬±gn¡IMt»Tn~E¶kau`±Lï@Ão·QY efQÈkpx¦]p²®bÀM VIĂNÊjä¶RnÂ´¸¾xì¬Ì¨®tĊlzp@®@ ^xäÔ²ðĜGÒS~TÎK¶R^BcQD½lEjrxtEAª}Böâ¢«nt~Dhv¦^BzļD¢e~HàrJÔ«FwĠǃA£g¯mw^`YyÃqJeɕÛśǁɱсuMÝ\\õURs@ßmsykE¿]ÃwgH]SO£Uu]cM¥©"],
                "encodeOffsets": [[-74644, 10685]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Nariño"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@`mq¡¢čÛImOt§Rą¤wT£Sn¡[ǁZ_²kUXÑb_Mud¥oiÃdwóJqjïlOvčgAĔZzF[lÓSµfs¼h`P¥aaRUao¥KÉ\\jOblÅngyPgW}\\a¸ãº~sjèÁagkkPi~u]uRÍÄ³E~cfc|µpgPf½kxPratoWezË_z°¸¸tÈM]Vy¨lÈLHcc|CôOZAE`~qtoFyEÒàMªhv@rt ¨ÞT®øìÜh¬Tl~jAZbe¬ºPh¡hkvªhngnXnWyh[np{r¾XrhZXB¥UPĉ·@n©ZOIdKjm [h¬^G¦ÈOzüÖ}B³Dÿm³ ypwMdY{yQDssuµVç®oKrZ²N_¸T\\vf¶A^Å ufmm£L}ßVÏ[X"],
                "encodeOffsets": [[-78766, 1378]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Putumayo"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@ĶÂ¨Z¤nrn~WÎÃDÏWkb­vo@½mÙl¡ªQZsªpbÊFcXNâT¦kîN`^Æ«QrS\\¶³tjybrr¶ChxYRyÒ_E}zmM´qdxAĊ£c]foNueP~HgI~jCtq¬SÝ}[T£p]goAs²V¶DôU¤LMjqi¢£Vuv_Qî{v~XtCbªlV¬PxPSÁ¾Nr|WDrSdtxTPXt}zUĒi{ґŹKORJdiCaxFKwVJlÃ]³HÖ}Y_{Yg~yOqvcLy­Vv£IWfoAbuMVa@¥SWkNħU½¡¸Í¾iLOaB¼sJmzsFcS¢le_«|{DAp©Pk¡N]Bŗq^[e±Cyl¥{mDiěvyl¥suZKpEuxÇOIqvMxpjc¦Nva`WÒV±lY` ǂ\\m¢T¤S £xQĆs¨PJnĎÜ¢¡nr_ "],
                "encodeOffsets": [[-78766, 1378]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Santander"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@rT¢¨v´°`ElyzHc¬R@jvªj¦XÖÊ~²À¾Ü ¤®WÀsvmÒ^@¶VvJÒ\\iØbvxðZæc¦I¤[CqvGÚ`bT¯rp­Kg H§áÉX]yİ@ziE¦»Py¬¿¨Qz\\pØH]K~¦ÀEeÚ@\\b®WSk\\¿_Muv«}N·uLslCeæ§FZx\\@xXj]T©Cat}mR±HëµÙLUyRYoyU[VühgpqRstUfyjCVYYXÕkmYbÍo{@wÁC{m³·O¯oqbkUe±MwZgkCk¡WpiiUËYWy`Cr~zÀJàF¸éaQ{Ã}gÏsg~{ßC¥o{goPcrW©­pKRGacÌAf­Zo²Bk·ÁmY\\vWpXtC²bkOi@ ZKknLhcJ^"],
                "encodeOffsets": [[-76309, 6427]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Tolima"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@V¡Kq^sa^³Ye@[uVeub_ùfaJs[Mo[VGji}dYKÃf{ésówcOwXqÔhzCZqW¨´`Nz®lh«fWL½kGl[Ë[H]k¿ÉÍ«myTUiOOZ¡Ã·©÷Y¯VMxBÈnDI÷FÅefuOU­v¿N½¯£PO{YLcwyHo_XWËnoµ¥I_Kw·Cwvû¨[[EwJÂouLOPe°^U|VhgªXjLÀ´ļQÞQ|f®ºCx¨¦Dlx® Q²HzªðtİB X¤®²H`hEG¢T°ÐyÊAA|{xNvÆ²_Pr[¸zTz¤bnEpÂ_zhĪ{"],
                "encodeOffsets": [[-76527, 5414]]
            }
        }, {
            "type": "Feature",
            "properties": {"name": "Vichada"},
            "geometry": {
                "type": "Polygon",
                "coordinates": ["@@ÀĪ¸èl¦âÌ\\W¾HÖvxÖ°¬lÐ¼ÀªVvfèzìĊJ~¶¨dÞĚ|pd¤fjxbêcðGĂNèptvdÌZp¡ØUt|tFzä~vwFg¢SªZTIhgE jF]jrÊXÔCmDÌZfFUpoÄV¬jt¸ V{¢L´SB{¦Cdkwdm{sQÕzk`£©}_m¡w¥}oçU[KñdåjicwO·£{Ó©±DpoE±y«¤é_E§YqDÅXE«Rãw¿Lécdg_XmK}rcK^uA_s^KrÃuûó«uq³w[ ÃvoKzO{P[}¢YsqZ½@fULUu\\¯@}^{TsoUeIyY{«QN[uUkhwGA{q@mebgkbÁ³afsaqK}IÑtLµCcOOkáT[kH­½mb{@p_Y±_|£{Ew¨mA·Q][qd¹q[m^kq§SwkOVo¤WO©_ajqgc}G{y]CxQ^¡F­ß}B§gčuImé´syraqCi_gZigRmy][oy^OuTF}[et¡Em@Њ@Ȅ@Ҟ@Ј@ˢ"],
                "encodeOffsets": [[-72783, 5043]]
            }
        }],
        "UTF8Encoding": true
    });
}));