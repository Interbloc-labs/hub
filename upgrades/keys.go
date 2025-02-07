package upgrades

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Name = "upgrade-4"
)

var (
	proposal14Voters = []string{
		"sent103l025fw2x7d8k8px7d07vu667x0h5p5hjxw6y",
		"sent10azrzn45u47n7m07vs4tntsrq5wchgc9w40ch4",
		"sent10h8qcursghmc94h3fgm0edyjrptupuzgskpnys",
		"sent10wrjthp576ymgwycrp37cuf0v0y8yf5pwpnaes",
		"sent120v78s0lf7umdmwnhct202ke37r9ymnek9tta3",
		"sent125hwc6q8zt7pfyqeew4fgwrq6l9u9l6xcak89w",
		"sent128xpn50a0ncum8w0cf5lcy9vu9793dyu3n60ev",
		"sent12g9tv7kzg06nx28szkvp7ud5jywpqv9samcky3",
		"sent12gmpp2fp2axr7r92w2na6ucq04svvka60wyn2l",
		"sent12k6x360zenzummgfqkats6avq7x3az0s07e75m",
		"sent12ysnawwxejaqu9rz7a4vpp7pmghqvmcyezj08p",
		"sent1302djxkzje9zkt0ry88hm0zq6ndswmuu57ta7y",
		"sent136gtx8m238907x02rnhxpwkwvwn0hpvhv8e3dy",
		"sent13ym9xeznvcpwyr9ytte3cf433zdkygq86nkh0v",
		"sent1488kv08va80uyw3rqm03dxapmgak6mx49ds270",
		"sent14dmyj9hm4klut74fgp3fxht77re46d74dg47n4",
		"sent14nnv9lzp0zkcfxgudmyzqtapge9k6z2ukkv6js",
		"sent14qvx8d42ulqhn2axvgee2cg5szxhppk0k5z5sl",
		"sent14t4czkut70d5dr7k55xfh337kka9n6em9pgjld",
		"sent14thklj2cgm830kv2l7n2cr02u7s83cl859karx",
		"sent14um3phqv5m3qsqz3fdtg5qljkrxs5je3jnvl23",
		"sent14uutcxkwe0rws08j74v8vl9f9z0qz3kc3t3sfq",
		"sent14uzu9dcdqaw0fnnxtjk0uy92c3rv3ae6lpktfj",
		"sent152mz474642sw4fclyg0h03eldwtemjycws3txp",
		"sent157nlexej8r55s3jafy3dp4x5l7fs0gxr7auml3",
		"sent157s6rkffq8vaqlmta26wr9c4a2fvcyd2xpjfx3",
		"sent15elks2ur8hdnj89wfeeem5y0wxjq7zn7w7ltn5",
		"sent15f2fn2t34y4sh475mx2lpsxctlptpcfgjd372q",
		"sent15ggj3hch4wfrzwehchpavjzqhcl76mn65zenw0",
		"sent15pdycya5fgxfmymfdrfdnsd874sh8w00z7xecx",
		"sent164zc5r2xufynsdywyv22eqdpav2lmgmam73sxx",
		"sent16ks39f8d86zhaus8xdppu2y0vge90j38zz8l9d",
		"sent16m5pj85dq9xsma000lugxrcqlfrz0xmnwjsutl",
		"sent16rmwfmxhc5swrr0s2qm9kq2egan2z9ny6xgvtf",
		"sent16vg279vhgempqyl8q5pr03fpqz7d7f05lvmh9g",
		"sent16x9fy0zkhmtq0f9rcmg3m9yryg3mmpel8wr575",
		"sent17czz0ftq5w52hzcgjmwzx82u0a4d5fgnptheuq",
		"sent17lxm4cvx6g3ya92psy83pz30g7uu0zknew59tm",
		"sent17t08gk6tehycwc3cza0myqu2kyxrf6k8e8w7nh",
		"sent184vnnwy3j4zmsjrml28y8heq47uhq8ar8k7apx",
		"sent1870tuajgsl4za7lwsgh0zf4pku3txy2d4h207x",
		"sent187wy2grtqlxpp6j3gspyaej3c58efargqw5g23",
		"sent188gynszvnt6fqr8h0ppudsltcff5yf7vat5gkk",
		"sent189g9war7xmrl68g4rxjgafjl0e99jnh97dz5rp",
		"sent189ym2ws9u8y6fk77dxe635kw9uyz8fr2000hdw",
		"sent18u2dtwkl85knxyp99ze38c2mvkt7ajcdsf5hg4",
		"sent18y66m60xn0wsp57ykagt3psg5fqzhv8mdks7z4",
		"sent192lefgwuxs9n9s6xp7mqxdw0wam5x54qewvt2q",
		"sent198u54gttp4ecwmg4putuvz379at7vxfn9nk2vf",
		"sent19dlpu8wk43eazrjpyk9068uey4s2vrmtjct4rn",
		"sent19h327payr2jr4qc2jvr55jfm9p76a07u83em02",
		"sent19hyjtyc8z06tf0pj67q63z955c57fq0d6fhjwu",
		"sent19j8nt48al76j92np9p84awnkdwukg95gv4zgqw",
		"sent19ztrw2k2gnzushs45fra0zteztm4dlu3qqj2jt",
		"sent1aksct9rrjujk9fnanfutesgmdat2naxkx7pvhj",
		"sent1awe6ky4pdvgu8rv9nmecd0rtcjtkpjrx88mrja",
		"sent1ax7nplm3jegg4z5ayfjkamrfh4a4e8pw0djjt5",
		"sent1azl9zpy6qrmxya94rud3gszyyrcvpj6pz7epkr",
		"sent1azvc3wmszj7nu9vaseu0u8s7v7w284adx4qt3e",
		"sent1c7tqfasq8smglwxzh8yaqs6zu43dh3vl7paajw",
		"sent1ceeqym2q878ek9337a5m7dp724cap00nyen95j",
		"sent1cy7kzvzs9f3vnfwrry4nvq44htr58rd7e836gr",
		"sent1d0ll6q9k4604kqkr8kpd4p9nkgtartv9vp6qs4",
		"sent1d3957x4h9g4h54v900vcsyfwlf5qnrcadvlmf4",
		"sent1deause75qhyz0mugkpwqfptczclmk9hf5xefyg",
		"sent1dfsd7aygvrmlkl0ymexjc30acgdy0mk07vxh6s",
		"sent1dq3qglnwry6u3pafx6r6u74lg5wkvyjj8gc9sm",
		"sent1dufaupc9cdprj0zp9pp6cgtfgqhndaefuagyqv",
		"sent1dzf04eymc4sf87evqxc87n8x6w75f867zr64en",
		"sent1e2h92rpz4vwthzss79k0hta8mgtk65h43k09p9",
		"sent1e2qur2pmzn5jtderg3p7hwck53ert72tl9elcu",
		"sent1ed8eahr4mnnz296gjetc3lhj2qlydjkqvqu8wk",
		"sent1ejvyzku2upg0q5a6n9nd9cwa29a830heqs4azl",
		"sent1euyp0d72gyuz2fnlws05s8vwushvca5zc5pvgq",
		"sent1ey0zmwryem278r7rfhj90uxfuq4ndl9jpxj88j",
		"sent1f4kmcewkr5jatxu7gm2em6hqst2ujg8ehl647r",
		"sent1f6feajvfzskwkm3jgz62la3f8l3nr8u29w5cpz",
		"sent1fcamvvfm3s6hvvnvkj8tkteacdyg3xcekp2qav",
		"sent1fxh8e0e0nl72xkn6sxtxe5qlky3mfjry9uhwhz",
		"sent1fy00ymcg0hl2d09422usyn40rmeyk09rx0etf4",
		"sent1g2mjkrugdjr0h75cwg9ml5r8ynltd4z9wulf5z",
		"sent1g5urq7wqff2asddv2jwefadhhnsre2ssytrry4",
		"sent1gtaegw4m2n7ww538fn3chnmnda5mqe73da5njl",
		"sent1gx5cx6prh66wkqtpmssd84q5dmvuw8cv7mqppf",
		"sent1h4w509ewf0wh6m3d9fqyak9z45m9lrqsh5xnrw",
		"sent1h63qg03qgfvcuxtzpwckhj0mehnm36dang7ue9",
		"sent1h79gn3u2hfeyvqz4npct574zr3vdqspefu2xcs",
		"sent1h9jq9l0ryec5newjttf84pp2fd52k823j2yt4h",
		"sent1hlxsavlny5ax6uwggr49qaw57cxnhcldar8sv9",
		"sent1hq642e0cjsg3a0lqtmn6gqrp280q7xe6lvvh4k",
		"sent1hqfslmsvg4zj2lwa238qlyqjvyntuechs2rkex",
		"sent1hyl79mayry4p2cees3dzc3qnwhdydjf238d3sq",
		"sent1j3gxknqug5lgfvnzp8a3f8xw7vnxdwrfdlsagu",
		"sent1j5l0tg7nxjj84hnkyyywph2avulyzysh562p8t",
		"sent1jev7zyn0zkln3gtsuprgwwxmrs6shsuxlxhqa4",
		"sent1jmjjga3j2qymr2v75x3p92wwmv0x6ch6fx3rru",
		"sent1jnpevf3zg3wc3vyqssxu96lz899g00sk9622e4",
		"sent1jxd29h6ra4taagzxp7k05f56u50h9r8zvduep5",
		"sent1jxy4tyx7j99jjaw203r8qxtgm73l9nkyuapmgn",
		"sent1k8f0jnzkq05zxeaugll8srsr8t2dyfj6zz7aa0",
		"sent1k9au9alnepaejtmn8zhl64a3uexfdf47l5ehxd",
		"sent1l2h0y53f0ehg23vps9ljmawlegzjd5r8wf8j7h",
		"sent1lejr8qt9ehyy7fxkkjc0h45mkjzr4dkzpreajh",
		"sent1lq0lyqnullv5aayaes29n8mnl4u50a7ymktw9t",
		"sent1m20782uxwdwwmdqqtaf078le7zfdnty8ezmqc2",
		"sent1m8du4mehzfl0y8ps72a6fzazjgg70af0y22hu2",
		"sent1memcf6sqd00h6yzf0e8e8e707tjmszpdmw93ym",
		"sent1myl82tug0sp4z644zxhnzkwfexl4h49s5gcxpj",
		"sent1n35de3qa3nvptxva5m0j6qucrjq0jnwadyk7hv",
		"sent1n5dufyjzwapk8ey7ysal48j6r4pfncamzg73hp",
		"sent1n6dgfnp9dmqlxrwv09gzv06x2ym8lcea639l5y",
		"sent1nchc75uswex4v5sa7fvprqpvr7jelnd3l6u0p5",
		"sent1nzys3gp94uyyeqk98qlhugzxgpvnv0njqqvnak",
		"sent1p6ly5kpq8j00fnhm9h3sltfxtr08x4qqqskm4r",
		"sent1pdma6s9lffu7m4h587yvnlkzzlq0vx2a5hnscz",
		"sent1pdsepwmf42mw53tcq70fe2xjfshgkrgne24dau",
		"sent1phcmh6rt32n05yfujxfphvv4y2a06nfz0ed2fv",
		"sent1prcmpmwggaxt8v8csvcpz9vc6dnnrr5qvc2c7l",
		"sent1ptsqy60scpjp3hyvhwn5u2jdymte8mp2duvxfj",
		"sent1qenepajhe7kxuvn3mrtustkuchhar5xfeyjla3",
		"sent1qmcsvf4nxctgxdl8fnyv2eldxv597nlxl7jwan",
		"sent1qnlxnneluedfg3lusv8r36t5c0s4r2p07l46c6",
		"sent1qqmhmrx9sd75kkfyyg9uncm326l8kclcztrg9f",
		"sent1rcvsjxmf9nnq9vqncrj57n4d7p8uw0utth56jy",
		"sent1rknrapqtccrj8qflfaywajg2yj2dd540evwmn6",
		"sent1rtw40svr7qlyxgwza95gkqnm8dpfmh7wj6e8kh",
		"sent1ruw9jyp5yax2ztkt05cx20x6zre3kp6cavzxcn",
		"sent1sc9h4dlftwqv82utj52xfgrv0xevuwngxka80d",
		"sent1sghqztvgjv3ug6ltpvvwwyhpt35u7hu8tv95mk",
		"sent1sv4hvhswhj4hfkq52mwa9g2v4z87qsjlk6pvgp",
		"sent1svw5zf0vqkdzwyuf4xage0f7nyhjw622ux6un2",
		"sent1t6exp5dc29ljwy93vlp3rmlgs0dvxhas0g3k85",
		"sent1tczkh46ey6rqk4mu5c9eujcs0ep5fdwphaaxav",
		"sent1tg35xn6jf3h65q40eaymqy45gwtu0rnft09k28",
		"sent1tlkhrt874den98e0jed4qsamhf2r85pszsr3r4",
		"sent1tpkk5sj5sek2h2k0a0mqd8lnk4y96vj4wvdhd3",
		"sent1tqpec8anmq498cwn4awst99usaflruha8ccsyl",
		"sent1tvcd8welups8wwgj9jxqpt9027eyzatlw6luyj",
		"sent1tyhz04rp93x2uvat3mpd67pkdgsgdrymkn9qpx",
		"sent1u79unmyzd3ggcdpayz4u6dhggx3tc6udsnp7eh",
		"sent1ukqrfz2ygwutz280gg3z2ksn7ukpxzp9ejwyxx",
		"sent1uqmgejdw24rgx6f4k9fqxf3ynwdr6ykp4n95km",
		"sent1uujlw6n4ha7sn0l7294swxrsjnq9xc34z4lpjg",
		"sent1uxukrn7amaent2ayjtclcvs8nhkjcs59k7wvsd",
		"sent1uzss2amzq34ryjvr7rt2kltsvxq20ak2y5hl6h",
		"sent1v9w49rtdpkwxeje8xnx9pnf9z9hyr2e0hyz87l",
		"sent1vf68rkgk502k4ja6l2mq9t8pcdrp5qpzqym029",
		"sent1vw8lk008zj2xgcnnmkpscj8e6fs0g4zu7zgyuh",
		"sent1vxhwam8vhntl3qslyxp663vr7v50j5h8j8any3",
		"sent1w5svr8h6lc5xckaz9vwcuqhwuvu5gxh50ceaa9",
		"sent1w8800khamp94n7htclskt0hdlxe9rn8cchmraj",
		"sent1w8cs54tjx5yc9hzp0e2pfmde905pr670qw7u06",
		"sent1wvjv9ujz5vy25ps8p2cal4f2vf2cmkladmzwhx",
		"sent1x62vpe4eqtmfwfv2j94hnwv40nryl36kv7nxp0",
		"sent1x8pt3hdf47fshryalrkvfcu8hzw0tly68lfxje",
		"sent1xjfgaw24t5yl094l72r9wdz03tzhpdajalspm2",
		"sent1xjn96jm0jjuhclp3x9r5kuncazaegk625jhj48",
		"sent1xrve72ecqqkmnsp76e5n60j764rqrwfavetvpy",
		"sent1xtlvlce8uxfzz092gtntn5xke3mu8hhufs3e0y",
		"sent1xwpgxsf39qv5zt2vva4fgsvrgmpxzh5ad2x3y2",
		"sent1y5ny4xt3afkdmcwa42j3c36f39547th80ncdzu",
		"sent1yaanlvrpys3l8dktea726grydtnad8we2rdvm4",
		"sent1yjtwqdfxul9ulfavpdgdzgghp3ga2ny2wf538x",
		"sent1ynlzcf7x0uxwk35u6kvtpsd23sfyake03zqct4",
		"sent1ypjcqq02nfvxgckdmafh9z5deg9lr89w0dg4ee",
		"sent1yvk507d2cthqmxvuw4408reuyzd54m3nkl0mkm",
		"sent1yxp82qlay8768867clrg64kje9dyv4y6cvu4xh",
		"sent1z2p8yxg9282pvyehht725l48w8qkpc8uwfen2g",
		"sent1z6pqtr53tkde6xxypt7mhw8gse8ld4dujrzrq5",
		"sent1z9z5t7fcn9257prq2rq7tua9sghxyat2yg6gka",
		"sent1ze0h7kf8044ekqkcmwcxnhrjcevwk0tzm4tjk0",
		"sent1zh66emq9l3qxj4ezhevl8wxd9qzkprqgmejefk",
		"sent1zjyx74rdw5ltge6vrwfzc5j0va6zdz93s9rkq4",
		"sent1zlm0u44mnf90wr9450cfqm3yfzxwv5zpc05fdp",
		"sent1zpgeqr7r8rljfuuf4492nvy3hkh6d0zwx0whf3",
		"sent1zt8ld4q488wrz3jpqm75mumh59hvp94j65p77h",
		"sent1ztz8t88xgqvueap24ufpfun5mxc3jhmct654sz",
		"sent1zzuttewnk3keg6m3xf2n3wxd2h7wkvj7kqqq7r",
	}
	proposal14BonusRate     = sdk.NewDecWithPrec(6, 2)
	proposal14LockPeriod    = (1.25 * 60 * 60 * 24 * 365.25) * time.Second
	proposal14VestingPeriod = (0.5 * 60 * 60 * 24 * 365.25) * time.Second

	proposal15Voters = []string{
		"sent106v0dgp92mwgerjph7aw84wwkutzuq4zqzsygc",
		"sent10qruk92y2ukzrlyqyqrqmszdxrejq6h6cn77tl",
		"sent13u9fk0juue6uyf9t6vrsn6c2czt30dacf24tgm",
		"sent14t3mz46apz934m59juprhygxz9qx09a9kh657n",
		"sent158wxl63rferzpwc6kl0fuphsw3njeesvp9v0y4",
		"sent15a2pxnp9yqxff0cxz2r7m4xvxknceur7wsvjq9",
		"sent15xtvkqrgd40fsfugf0wag5hps9ctxy6qj4jdxu",
		"sent16hfcr3xfr0c0jt3lmqg5qfczrtltvx6svc0xnn",
		"sent16s3nqpvl6jypaxg5j2u68kejf5fuegtxce0d95",
		"sent17e7f9rh93jj9vvq38nzjygaf9tnq5dy8xs64xw",
		"sent1adc0sjsw3azqxdlngq5wzlr93zxxsjtxhac39d",
		"sent1aq56hjusnaf8nz4knays2yj7u789nvqxdqatvy",
		"sent1ccpmdkndqsur6anr4g0ys5q7l22hmr66d857tn",
		"sent1dee0v47ducs0g2rm5e2urx7ruqq43zxa60tlea",
		"sent1drvdx4g9k6p73lt68ld3y3aqqyd7ylwc4hzqm0",
		"sent1g4spcxptk8akl5assv7k0mrdwuahwfjywyjawy",
		"sent1hsxezlay09t8ty6maue78gr8rw7m86l3d4qr5c",
		"sent1jh467gqc09r6eza28ewxy5f4yv9gjx85jzta9c",
		"sent1kd8vrc7xxtt4dw7ghv36kez4vzuqryqvf8nj6r",
		"sent1kqc683nnl6u0ll6kdpymvzgs4f8p7huvngsp5n",
		"sent1lff2vtk36hxrlm35caeqss3vrn3kwz504aal69",
		"sent1lh78a0p9p744uas9kryx4gwq3zggny074sk36u",
		"sent1muu9z9l95dt3skecrpck9x7v044rs6hqp44qgf",
		"sent1n0h54c882xqalqeks8vy45ftl2w74naq9tg3fc",
		"sent1naepmtr5cfse2jlt2c95dcfy03xljtrd9mhztr",
		"sent1nczvm4k9k7a00u4t8x9kcxqh2x6fjfp55tgx96",
		"sent1qggq0f0pkk45sqmwme2crqsvt4stuymtknlnnl",
		"sent1qlljvx0ggg0qmg3ffz4f3hd0tsd259n4fkqjqt",
		"sent1qudupvsq3rekpwnwqxvqhcre7xxh6kaht0tjgv",
		"sent1r8fysr2z6rptpc2j4c8t9kmen20njsk4mcdt4e",
		"sent1rf5hprmts0h50ww7mktkly9pw9sdkyt6x8z9me",
		"sent1uyts2z7dc52aptskxcen8zh839lglxphynn9gr",
		"sent1vss9v2rq0ggnnyn4ypu3dpu8djuuhveuzdhydc",
		"sent1vzfkvyl07xcmt55spdqehz89mgue92kpfncqk8",
		"sent1w6ha70ghjp4rg57cvtws6k5jnz8t85a0pp8nj8",
		"sent1wj204t7hz2anaffvnwvkvuyexclcejdejshwt6",
		"sent1zmyx0tegzetxfvlutf09fagmugz2r044cgpk83",
	}
	proposal15BonusRate     = sdk.NewDecWithPrec(12, 2)
	proposal15LockPeriod    = (2.5 * 60 * 60 * 24 * 365.25) * time.Second
	proposal15VestingPeriod = (1 * 60 * 60 * 24 * 365.25) * time.Second

	proposal16Voters = []string{
		"sent100fjd4fgvavlevh44wqtfl8e46ryyzmhfxa5rp",
		"sent10cqu7a6jyh86a595d30p2k0wpdnl49st3wltcc",
		"sent10e2v6zzv5kz5fky0pkdwssg4elfglgccpr2x9g",
		"sent10v3veqwsr7jvc6gjhtjsflutq9yzyxfqeqr202",
		"sent10zhjwdmn0c6v0kn80t5l0v2dnxak6m5j9hc957",
		"sent120d7u66wk2w6rvjtsrdjj3c6pfzkm8hxn37qhz",
		"sent13f8xsnlf0ejwsky259gyqv7yrgtr2p6lyt98lh",
		"sent13sa5af94t0y69j86dfhkd27yunrygv3pw5chjg",
		"sent13sc66ae36a8m6zcqyg8h63t9syjeg5a7rh82gn",
		"sent14dvr4kzzru0xdhflz566xzfprf90hdw830fw0y",
		"sent15jc2epg2kxle5vlan7k44txnceyf6muf6ftz93",
		"sent15md4wkfqux35ak5rh2409yk3gx8l4y5shxu5a0",
		"sent15p4gfmq4p9qqt6yhsd3whqu8qtej8xgdk8yktg",
		"sent164rshrwr7h035m4hy7yvvlg2gav8yn0ae2dha2",
		"sent1697eu6zsqnghmac28a2d30y3ffxehpztze76h0",
		"sent16a6feh58x3v3cty8vfzzj99q78upux70zjznn0",
		"sent16l6pz9cg4fa7nsfmjdj80n94gae7nqr9d4t5jj",
		"sent16lnqygvcmmd2z63dzuyx80pgyzdxnrsn0m5mt3",
		"sent16r7w8vpsc4gtws2wqelk2sz53ra40het0d0p9m",
		"sent16wfx5q6j7uzqh2vz8g0kx7r7ka4qegrx3gcgth",
		"sent16xgz5gdtnxp5t9ecychk9eft0alyvxs8ulyjvh",
		"sent170rshs95ty3xj5j9qggttcndnwrg0ylay3j3qg",
		"sent17asmacqf9wn3ghsmjvtljez7zweutdyh45cpjw",
		"sent17c6syh2twxa3tqwlrg945kj7exvz825xaq8k0s",
		"sent17xlvkaj69qnlnp8ny53hlapxxgmfseut32045s",
		"sent18ymvc62csr5qkmht6s5uy20xd604h26663g94z",
		"sent190kem4hm3vkljaeyyldlmnjnzw5l7t9h644mfm",
		"sent19yrexd5s8j9eg8lvxtah2y04np8a6lcvfqc8zq",
		"sent1a69czspmvqxpf0g6pqheleadv5xk7m0v2hms5h",
		"sent1a9fzkadrn347lzwr2wfzn9mvg4a5m0qul0l9nz",
		"sent1a9h9wjwpaxhuu9exfjfrky6rwzhtungdc36w20",
		"sent1aqrm0jf833qaqnf2fepd5n38tfwy549xevw484",
		"sent1c0c69p362467emvcwcz84n5lk6grqe04cnu48m",
		"sent1czdf2ds33f57thmu4v5negsqz4jmlu2um5u378",
		"sent1d8mq46wt2yxsgwrmh6hhfgycl0537w8gsq3lyp",
		"sent1dapuw6zg78z609kqk0vyxr95c4k8a0tx3q6eqy",
		"sent1dlyccvmvcewd2mahhs5n69n3ves4ldtlhmskhj",
		"sent1ffvhvyrntp4l99xqgn2vvly6v4kn2xlh4qetps",
		"sent1ftp0m225xqhhzeju532s0lwzk4xyr2svwk9pu4",
		"sent1g96zjr90sew2tsewp8q7fx8utg00efm3akxl2n",
		"sent1gh87xjzr70saaqhjmw3p0xc9wkpy7f7yfdpvx9",
		"sent1gqyavuuyv5mhj9jgk2sxlwxfcngxctf9rsx2vc",
		"sent1gszx0trncayqq3q7qwqquzqrfq3ee8e2z7h45y",
		"sent1gzln7fwa3md45al9e4cg75q77j8zy3gvq6za7u",
		"sent1h8n9h6g8ugwwx8qr9jp834akv3wrlmhd43ma7r",
		"sent1hjtnpfgykv68nll0yl7xana45n9e9rxzy3maqm",
		"sent1jkl7tkgaq5vky4dhafyrvqzjtxhtcvj73v65gw",
		"sent1jl82qy8645dpmrcxe64wfxtq8rygyxlr4aspd0",
		"sent1jshcg82q4p048gduml28hdudnuqs87nes8yh38",
		"sent1k5w5fpu3qt4va27vsarddnruumqqprs8dnp8qd",
		"sent1kgkp2p6x73y5j8mlkjxhxgtf3e3sgx5gh6v925",
		"sent1kxykphkk82x4pk5nzq2kq5f2r6n69k73peutee",
		"sent1kyd3kmm87fstesqe7970w6g6mt6j4eahsy6xmt",
		"sent1l7t2xupwtv287mds7xwvy7yas59n50l4855jwm",
		"sent1lc6sv2d5qd7dsjcchyk50hzu0c0s0cdmza6ctr",
		"sent1lltxfhkqyhd835ha3j5690wc4s0canjwalc2j7",
		"sent1lz24q5pvnnwnwxlajhdsqjpn24scszjpeut2aw",
		"sent1m2xeyv99hs58g2p7gzcem3eayack9r2rtt7eq0",
		"sent1maxml5v7veytlmucp2al9hwffge7p9em6g6ess",
		"sent1mhd932nfryqkfafdqt7qta0dgc4kr56y54t5mn",
		"sent1n00fcnszy0qymnw392tw2awmf3q687xqzetkd7",
		"sent1ncpnpe5n94582kvcn4574dhxl4fk30srahs4af",
		"sent1ndp69jt2x6ej9may75ufynx7c9krnmsvwq7drl",
		"sent1nuv7tx6u6vp5tm7ewymqwfvugvyk4937czl8wu",
		"sent1nvr5n54cewuvefqv0v3ckgrz5kgy4sg7ckezcf",
		"sent1p3z4l0gcnr7yywxcuk5hccq64srpqajwm66erz",
		"sent1pdt3hl4ef0g9v80xudh6m80a7ryqnzmxg8nk3q",
		"sent1pfdg20wh4vfcfaxn72qegqxlq7cvllx30ja4cg",
		"sent1pyccy8msrhy7epfrs06s56xja37hwnwnl6c4g9",
		"sent1qn6clt2v2mdr4yufcgdtreq7a8ra0pn6vvez3g",
		"sent1qthhthjy68k4z8cjml95zq36jfsdxj83aktrun",
		"sent1qvtxtfmjakkqqhxrtz4r04rqyp9zwwdtgr22ay",
		"sent1rmp550cf6cwtk5w4phgq0vkyjjf3rya6pukcl9",
		"sent1s5c4wgq4uft4lkzlmmasssa5j3qcegx48p4l0c",
		"sent1s7p8qqe7ucs23nmnld3p746hsk25aujtxr6jqe",
		"sent1sexg35ajngt6hzc7fep8slj5nmsdqz685de7df",
		"sent1swqcn946mvgvfpup84msgjs96kx6yejft8vk5f",
		"sent1t97tmw74c5870chqp2pyd29hkn559mdk6zezu4",
		"sent1tkcdkh23du72nm02dq8w29f52lruwdc6a3v0m4",
		"sent1u7938vmh4y9ng7n0cfvkz3nj6gv7awyz27cpus",
		"sent1uh2gxycd8zf3vharxfekhhu0jddstw4xxlq767",
		"sent1uh7qk3mxa0hd4f3rr55rq6jty8szf3hesqx2xt",
		"sent1uvay4nxlvy4uud25dcvdmruhl4n0zg0dh5xej9",
		"sent1v59s5mpfq6fsm7apjzt9mk8h4ryccmhckzj5zm",
		"sent1vdqlc2xqd87vg4l7jq4ca9vwe4890slqqls0ku",
		"sent1vfm9rpggu5h44l627l4yvzdvnn3y0tdl6js9ef",
		"sent1vrg6g8jncpyjn03f6gqnks27adjd2c2xe30clz",
		"sent1vv8kmwrs24j5emzw8dp7k8satgea62l7knegd7",
		"sent1vyrp5ma6d2kzjww4crl4juefw79kstu4ry3xa7",
		"sent1wr64j7238sk89qtvxpauzwzwmppp5yz8s4erjq",
		"sent1wsgte9eamcchp9p77w7ryzyxys7aswv2jrsp8q",
		"sent1wzgqv5dncj5aa0yqd32h2q8kd949kzpwmgxsq3",
		"sent1x03scqem5gthxyhfefhsgl9a3y24k2lnq549ws",
		"sent1x4nsrhgak74fne0gv66alfje64cqgfs8vsyusc",
		"sent1xf4n5alscmyuhuyv5elnpns38lf64u8ytezu2x",
		"sent1xhjjrwx30qrwxsksgwqjl2qez0z9mhwkrrwm2f",
		"sent1xn9sjdhgjjpppj5ugzqzs4pv0j2dx6y597k6dq",
		"sent1xvtyjhv7h5fuvdfwhr3ezq8nrtpwjn9vgdp2ur",
		"sent1y0madmkzkhqcytafmtetgmfh5u0ymxu2kufq4e",
		"sent1yg4vpe9zr2mdv7jq3e3rr506ahrl963v6ll22x",
		"sent1yu4yc8sqku2rq8k3ar9vqa577reetknkywz029",
		"sent1z5tq60s9zxu339j9pvdud4qeaued0xpt48y4g0",
		"sent1zw3xlh7v5fj6nr6mnsjuu2z3hv6wjqskaghsyt",
	}
	proposal16BonusRate     = sdk.NewDecWithPrec(18, 2)
	proposal16LockPeriod    = (5 * 60 * 60 * 24 * 365.25) * time.Second
	proposal16VestingPeriod = (1.5 * 60 * 60 * 24 * 365.25) * time.Second
)
