package adv6

import (
	"fmt"
	"strings"
)

const test = "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
const puzzleInput = "qpicundo\nfiqcdbkyuoz\n\nrahgpijvyfd\nbiwvrajyp\najbrvopeiyw\n\ncv\nv\nqwvo\nv\n\nftsvmcrbxn\nzkvdfsjxrablnpoutg\nnbxferitsv\nnvfhytbrwxs\n\nhdamqjxnipezgrwutc\nvmszeqacbiuwpokxn\n\nkushbeg\nfjarlmdhvspkbie\ncesznhbxkg\nbuekysxh\nzbkhse\n\nhrxgtupf\ngmhfxutje\n\ncnajtoldrhkmiqx\nonmqriltahxjkcd\njqimcxkrdoanhlt\n\nfhbtngdzq\nndbghqftz\nqfhdgtznb\nnbdgtfhzq\n\ndxeoijmbprfaug\nohspnculriwexbk\nzuviobjraeqxgp\n\noitgeazcqrkyvmphl\nbapolshcnvekdyrijz\n\nhxaiwmzpglyrunqkdfvjbsc\njsymwqnfzhbvrdxcupglika\nbqzsuwpkfjxnildehagvyrmc\n\nagkfyuvncr\nkvucyrfgan\ncankyfurgv\n\nqrikblojp\nwtvfscqiyh\ncfwynimq\nwuqxmegszi\n\negvjnhmwuszotipdc\nnzgvmjshpcuetiwod\ndqwpiyzmcaoxuhetvngs\n\nxae\nxaet\nxea\nxae\nexa\n\njizaflhpxdwumgs\nduaflpzwxmsgjhi\nwfsmilagxhdzujp\n\nijeyq\nezqji\nzqeboij\nqeijb\n\nbad\ndba\ndebta\n\nh\nx\nx\nx\n\nulfkv\nkuly\nkuly\nkjul\nklu\n\ngpr\ntkr\nrd\nr\ndr\n\nmlcybiwua\nowg\n\ncokrh\nbomzakugxtwveiyn\nosklqp\nofkq\ncpqohk\n\nhcmyonbsarvldkgfpw\neslkdgounrvbfmywh\nslkdhzvbfnymorgw\n\nefvugtbyzmdh\nfumqpbvyej\nypcmbivuesf\n\nmdgfxrjycuwepi\nsdtmejcyubgqlox\nxdmyecujkgzafr\n\narhevtgdomli\nidvraethmglo\nehivtolgrpwa\n\nzrjkgal\nrcjga\nrljgasp\nwrbhjnav\njslra\n\ngo\ngoxw\n\nnfycq\nfnyc\n\nm\nkf\nj\nd\nq\n\nwcx\nxfcqw\nqxwtu\noxkpiyg\n\nud\nqr\nsomtbxcv\ndiae\nij\n\nvftcjdgs\ncdtvfgsj\ngvfsdcjt\nvgfksjdct\ngtefcdjvms\n\nkhgdmzlcae\nzedcslqka\ntfdckeaxmzlhr\naezkdrlcs\nwycukdlvaezb\n\nhzakytmxgcqvunsbjoidwe\ntouhjicbexangskvzdwmq\nubdepshljqaivfnzgxwromctk\n\nhxgcwlvzujrsbe\nzxceuhvwgrnldbsjm\nrvczewuhsxjlgb\n\nrhiuambjzqn\nimqzjpburanh\nbrmziaqujnh\naqifozjubhrnm\n\nukrpajtcnlvysdhwf\nrnytlpwaicjusd\nljyarntxhsbwucpd\nlrjwscyvdntuap\n\nxwgtyqkf\nwvuinb\nzlmewri\n\nvajupxzg\nbpzaux\npxzua\nzpuax\n\nwmqd\ndqmw\ngdqwm\nwqmd\numqrdw\n\nwrfhtn\nrofwht\npwsrthf\n\nmanfclbyrdxjik\nbxflmyadknjcir\n\nsduzqohjtnkmi\nhjtnyimzqdsuk\nmjquhntdiksz\n\nofr\nqrl\n\nyhjzgdmslpn\nzygdpjlmsn\n\nailyqtfmzsopr\niazfmqslyrjhpt\ntlpvqgaiuzsnmwfbeyx\ndclofymkatspizhq\n\nclypfnqvimtuxwab\nmsogekqrzhjf\n\nhnvgmwrfj\nrhfjgmw\nigdfsrhj\nhjfewlrg\n\njulft\nfuijtqyc\nowfjskvtg\nituejqfc\nyuztqfdcirj\n\nrubfcosnikjzmyav\nfymjibrucozavkns\nuzanfkocrsyjbiv\nyrbvncsjuoamkzfi\nsfcjvyaqodezubnipkr\n\nqkgwlfhdy\nnbjkrsveifz\nfoypxtcak\n\ncjqwymzgbxilsednupfovk\nfgqevluowysidbnmxjkzp\nliekuxymdsvbzqjnofwgtp\n\nsgfnmoljbdihqaeyuwcx\nojuwiqhdlnzfcsegxymab\nwbftqhedrmikaluxsgpjncov\n\npaxhq\ncos\nfdb\n\nognxqwer\nqwgoxenr\ngwqroxen\n\nfsdgo\nkdocfslx\nosbdif\nhosdprf\ngsvofzd\n\noamvucfhpq\nhquvglbsw\nqrvhu\nuvjhbltdq\n\njfltodbmwxpcai\ndwfihozpbcatljxm\npbwotmcldjxaif\ntxkbipfwdjmlcao\najfmdxlipbtocws\n\nlkbehfsx\nhxlksfebp\n\nlm\nwhnogp\nkicyqtudszf\njn\n\ngcsewnmqdthv\ngsbucadrmz\nrocjzgsldm\n\nuprsfb\ngfqnmlbjtev\nfhwb\n\nawpv\na\na\n\nt\nrtn\nt\n\nyrfmxdnkel\nlmdzewkxnr\ndrempnxkl\n\nkyh\nyfh\nkyh\nyhk\n\nybqx\nxbyq\nxnbqoyi\n\nrkelsyhwmjqoaxgbndzicu\nswlgkjmbqhacioedvnur\ndqorcpgjmhieunklwabs\nuoijceghaqxlbnrwskdm\n\nxmd\nmxd\nxdm\nxmd\n\nunjqpocarmtgfikxzsdhwlyebv\ngwklmrhxdvpfbzyacueitqnso\n\nwfhlcpe\nfqechzwal\nhlwfec\nywbefclh\nhecwlf\n\nv\nv\ns\n\ngzhspvqbtiyl\nltqiyurkhb\ntbhlyrqi\n\nhepdy\nydephi\nedsrhpy\nedyph\npyhde\n\nupcjxiwvqmhafetsronkbl\natflkcnqumpohixbzsrjwve\nwlovimkcxftspajbquerhn\nikqrstefvolhdwcbmapjuxn\nipsecmgurbxqtonvwjfkhal\n\nnrhjbowi\nwiornbhv\nbnwhroil\n\nszfukhordxjb\nrfbunjtdsk\nndbjrkufzhso\nukrjbndfs\neukgdjmrfsyb\n\ngbzd\nvdz\n\noqfzdjk\nkedfj\nbuwinxdrjgyh\n\ngy\neuy\nywz\nyk\n\nqfglvert\neftjpsln\n\nujqwnmhszxcopivfgbe\nwgunfyqopmzesadcvxjbt\n\nbaiykdzpg\npayezibmdgk\nikyzgpbda\npbdgakziy\npkdbazigy\n\npojhgwqdycz\nyldpsuwco\n\nlqiox\nkhglw\n\niknzjfuqcx\nqnkcovzui\nzpnrqauywtdlme\ngqunsfvzo\n\nayhdk\nojybpl\nyvpb\n\nubmxehjto\nxbtmoujeh\nxmoejubkt\njuoetxbm\n\nkvcbe\neckbpv\nwceqvkbm\n\nhdojvxsbaiy\nelwbsq\n\nmrng\nwlz\n\npbiluamrejsownyfkgx\nodbpxnuizywltqvcs\n\nyhgoj\ninls\nvqzmeawutpf\ncdyjxr\n\nweviflqhaotpu\ncsybrknxgm\nzjmdkr\n\nwnp\nwpn\npwn\nwpn\n\ntrfmhxeociukg\necioxufhrmtk\nthfxmenokcuir\nikxehroumftc\n\njgxzko\njizkl\nxoukjzv\n\nrkwnxumvyigodejlfp\nokerwlypnvxdi\nqztwohyxkcenadvbpr\n\nxtisomdj\nuncyzkoesravi\n\nyrplsicjnox\nriopycxsjlun\nolyjcpsixnr\nnxiprjclytos\n\nlypqxvjrueastb\nlvpruqbatyx\nynvurplqabtdx\n\na\naq\nw\nuzb\n\nb\nb\nb\nb\n\nbdhtie\nbthedi\ndihebt\nbqihdet\ndbethi\n\nbmcfdeuhykarnxq\nbuxaynpetdkhmifqg\nyabdjmxqunkehof\n\nhbjle\npyr\n\nga\nga\nag\n\nhit\neft\n\nbls\nsl\nysl\nbsl\n\nqjtvniz\njit\niyjt\ntij\n\ncindevpkljyarzqwh\nzgtdpmsxnofbyiru\n\nphismdwurq\nximrwtlyzqh\nhfmjziqraw\n\nay\nay\neya\nya\n\nelnvry\nckiyeqnvml\nyvenlzrs\nvzneyl\nrevluny\n\njbqenozxgd\nxdjqgnoezb\nnxqgeozbjd\ngzeqobndxj\nbqdegnzoxj\n\nwfroqvpdjyu\njprvwouqfdy\nvoudwqypfjr\njwuqyfpdrov\npjdwovyufrq\n\nwpmbq\nbqmpw\nmphqwb\n\nq\nq\nq\ne\nq\n\nmpwuo\nopbwa\nowgp\n\nrbmveola\neoizbmv\nlqbaormve\n\nenqk\nkqn\nknq\n\nfdoajsxqhlinrywbgu\nhxnuolryqdajgbifs\noaduxbjnigqflrysh\n\nr\ni\ni\ni\n\ndngie\nigdne\ndjfrgiwenmk\ngdnie\ngidne\n\nompatvbu\npkmuxavbt\nszufycrimbapw\nvbuapm\nkabupm\n\nquonbdt\notuybq\nbotqu\nuqotb\nqubozt\n\nj\nj\nyj\nj\n\nduqnjebpv\nlquynhc\n\nhndkojuxyslvrm\novmndkwyhxuisjlr\nkjlyhsdguprvoxnm\nlsomkhyrjdxunv\ngudhsnlrjkxmvyo\n\nvtpamce\novzaryjmetpc\nncavpt\ntpcavy\nbwptvxdqhaikufc\n\no\no\no\no\n\njqluihoeak\nikeujqoalh\njoahkluqie\n\nfqrmvtxas\nzoiwqlmt\nxvqhgebtsm\ntpmeq\n\nfkaeotvzuwyxdjs\nxafmzeycqbshvdgkl\ntyxasdvkefz\nvxoaksyzfed\n\nrkfvodejlb\nqahc\n\nurabynscomx\nbvmaqynlecf\nkhjlyzcdigpbtanm\n\nlosbxwmapcdrzfyj\nujmlczgbwxyehkoq\njcgqylwbmzxo\n\ng\ng\n\nbvs\ns\ns\n\nbhjwfznyraolkce\nfbqjlrcwaokte\nbvlrwotfecdkja\nrwkflbxijougmepac\nkfwajczeoblrt\n\nlhdcfskwxaerjiotzvn\nzaiefvhslwrcjoxk\nzoxjibkalmshrevcup\n\negvpws\nepwgsv\nvcpsgem\nozvyetksrxagqp\nlpmgseuv\n\noupbzjesrfwnvklxtiygdc\niyjudbvzegtscownlrkfp\ngwfzvlynutkprosedbicj\nwdkhujnigopzbcystelvrf\ndgrlvbtfzinpkusjwceoy\n\nxupqorents\npensurtqxo\norensquxpt\nprsuxonteq\nsnxqrpotue\n\nobryetcjszpnuwgxilkd\nropejxinuskybcgwltz\nzuvrgmwftjnelqsxkpcyi\nnxkzhlpgsdyrwejtciu\n\nyzriohapbgs\nisprgzobahy\nriyabgohszp\n\nlxczkmearuhnf\nmxfnkczruhela\n\nkrcoivbedawtzuhq\nwnksyfjarde\n\nxsguj\nusxjg\ngsujx\ngjxus\nugsjex\n\nhy\nhsn\npz\nc\nn\n\nwdnzt\nwqzrdkbc\nnzwdtf\n\nh\nmh\nbh\n\nkqalzofbu\nwyncpf\nixf\nhf\n\ntihepcdnxg\nscfhe\ncefh\n\nzxlhgpuowimedyjktqrc\nzgcirhetqjmlpdokuywx\noqykmcupwridlxhztjge\n\nxvstwya\nvtywxas\nyvxtaws\n\ngsnij\ngpjshn\nogjsnlze\njzongs\nswljegdn\n\nteaxnu\natreuxn\naxlwtnu\npvukhadbjtxcn\n\nuvgz\nzsm\n\nxrjduek\ndjuiker\n\nmgbsqftv\nbgvsfmqt\nftsvmqgb\ntvgfbmcqs\nmbfsqvtg\n\nopjuwfskvclrxgtbnzemqhia\njsutkmcybpgowzaqxrnfhliev\nwmhbsitkclvzfxgjpurneaqo\n\nxblyaorpvdjqfehuszi\npyihjwadusfqbv\n\nlhtxfsockq\nhsoflktqcx\nlkfhtxocsq\nfqocxklhst\n\nzwt\nzhtbw\nwzitlb\nwtbz\nzqjtw\n\npnowugqirtlmzadsv\nuoqpznmldrtwjvsi\n\nmeavkwnigcbjplhd\nminwphevajckdtgl\n\nudqekwjaxm\nxujwqmkead\nemauxkdqwj\n\nsxtnmaechq\nqnesxcth\nntkxshecq\n\nfl\nyt\n\npnjmvxgd\nxeqdpmsvgnj\nnvjpxmgd\npxngjvdm\n\nezfswkgqrjnbmpvliycxua\nubzspxqgnejywkmirfv\npbigrvjshkxemnfuwyqz\n\nunc\ncnua\nrknluc\nzebcntou\nfunc\n\nucbipqmenafw\ncqiamwenufb\nbqveahpzcuniwf\nuatniwcdbqxjeyrof\nfaeincqbuw\n\nmkbtyagfurhszxcpqjweivld\nzygiwpcnuxqjdlmvfeohsb\n\nqtczamlod\nqovpdgfla\nldaqi\ndwnarbsheyxl\natvuljqocdf\n\ncjxpytmesknol\nesxcmpoyljtn\neosljxtcpknmy\nptsoyjmclbxnve\n\ngsmxqruzwhtb\ngwrthqmisbv\ntgbhqmrsw\nsqwmgrhvtb\nhqrbsmfwtg\n\nxdzjpkolmgchyatnbqv\nyiuprfwxnosme\n\naruewcdm\nnroei\n\nyxh\nxhy\nhyqx\n\na\na\na\na\na\n\ngkrsznedvcwxpymtu\nirqenbkaox\n\nkajrbfgdlmenh\nmflgandejrhkybx\nafbmqelnjkhgrd\n\nvfmns\nzonlxhfrvm\n\nzcjhsuikpbaqgwnxrvedmfyt\ndvpnabeuwzqhrsgfyioxjtck\nqjwdzgurbcsehtxvkfanyip\najivfcuzwhrnptxyqbdskeg\nxjwdfupbnzsqygheiarkvct\n\ns\nb\ni\nb\n\nwqoxgkjatmlcynhs\nmojxlyhcnatwsk\nclktjowxayhsmn\nsxonlyakmtjhcw\nmawkojnhlystcx\n\nhovsj\nujvso\njsvo\n\nwljeskzntydaqfg\nplxzvtwdrqosfckhnjb\n\naswnqfrcoe\noecaqf\nziqayfjum\navfrq\n\niortchugjyax\nmlxsbgzrvfaun\n\nispmrkfvyhczdotnjua\nkqcvmrasoblntihzxufjydp\ncptoryiszakdhjmuvfn\n\nqstwglyd\nqylwdstg\ntgsdqwyl\nwsdygtql\n\nshcbwqgu\nbdqmn\nvaesbq\nlrptqfijxy\nqkz\n\nsvoxw\nzxwbkvo\nowzxve\nxwogbnvzk\nvtwfqjiox\n\nwq\nrwq\nqw\n\nnrxyvf\nyfpnrx\nfnxpvyrj\nedwsaytzfhnxr\n\njdpfyteizvkc\nypjmnvlwtfdcezgk\n\nbcpwnkgsvtxdoryifh\ngfscykroibnxdhtv\nbgyjnckfsrziqvohxdtu\n\njxwqdnocstrzvhybguilemf\nbmcqdgrlznwsexfavhyioj\nwxamfqjrobyzicesdnplgvh\nngjixrfmsvdopqbecwlyzh\ngnfdcimovhbjxrszwyleq\n\nrpgdzmcneyjthsvflb\ntnmfyjrzbosldg\nndszyijtbfrlgm\ntzslbmrfoynqjdg\n\nrekgsioqlu\nunjlczqkbdvgahm\nqpxrtusklgyo\n\niwupkfsxnoqzcd\npvzwkmdejybqhtrcn\nzswapofucqgndk\nukqgpdlocnwz\n\nyxzpgn\nxzypng\nxyzpng\n\nibg\nigbf\nigb\nbig\n\nkrqgmdcbxeuvsaihj\njsiartydzuqhvowp\nqvbrdskexjauih\n\nbgaoqndvurhy\nursayhgpqbonc\n\nnrzwfeqdmui\nrdjmwnkzpq\n\naihrwbyjzt\nwcajbti\n\nsizktqxjyohd\nkhzqdtxosjyi\n\nqclmepht\nqdplucnb\n\nmaxvquihzfskdbglj\nwvisljfgbkrmqzca\nnafgikszlbqjvm\nkfmsbcgvizalqj\n\nub\nbeu\n\ndxhnelp\nmpfiselg\nlpexzoaytv\n\ngmuvyxln\nwkxoegvjhumdal\n\nx\nx\nx\nx\nx\n\nskvxydbqmzehnlcwo\nvcnkdequzxloghybwp\n\nuslq\nsluq\nqsul\nosqul\n\neuknogsl\njbvrnkgoslu\nnpmcutyihaogwsxqdfl\n\nqbfoiukvea\npgdzhc\n\nnjdlgbus\njudbcghtnf\n\ncalds\nlcsa\n\ndvcnlxhkt\nbvzhingdrykjsxo\nvkphfxdn\n\nbsferyztuikaq\nfsbayqukrtzeic\nebksiyzlqrfuta\n\nspmkwhd\nnhbz\nbqxclfh\nhjnli\n\nhsmqoplv\nhopqvsl\nolbhfvwusp\nhoqpslve\nsvoplnh\n\ntpnuwvzabcmjfiod\navwntjdbzcupfo\n\nbdvyenjktqcgawm\nrvgteckqmb\nkqrvmtfbgec\n\nfu\nsdc\njzrginhoalw\n\nhpwdfomy\nshwjzmpfikoydb\nyfwmhdopl\nyhwdmfpo\n\nq\nq\nm\nh\nm\n\nqg\ngq\nqg\n\nd\ndo\nd\n\npebzkhvxmicuwgoynrjdsf\ndchofjsmzrkgwxpyvqeiunb\n\nljfm\nmflj\nljmf\nmvjufl\nmjfly\n\njuvli\nluvji\nlviuj\nvliju\nlijvu\n\nqyokiecwgxsdpf\npwqfselocinxgdky\nyoksiqgwdecpfx\n\ny\nxpofmyg\nvyhz\ny\nyz\n\nrihacgdmbn\nabcmigdhnr\nignabdmchr\ngdhmirabnc\nrdhganmicb\n\nbizqcode\nzdp\ndpazjrt\nwzxd\n\ncmfihuwqnzbtglp\nitnwhuzlpgjmfqcdb\n\nut\nut\nut\nut\ntud\n\nklhufpyioetxngj\njxuhfpiyntkgeo\nhkefjnpgixuyto\nigknoteuxpfhjy\n\nuysoaqjzmgcrvkdnlebpwtx\nenlrwtxomaqjupvsczydkgb\nvdsziegptqljwabxnoruycmk\n\nnycf\nmcxo\n\nytivwfslq\nftvqli\nfvitquel\n\nnybhaiupejdxkgs\nxyajdsbuekinp\nayjpbnuxksedcw\npeynrkadsuzjxb\nnugdjbypkxashe\n\nnvdaqfkbyomhxptiswugrzce\nxzkhfsjotigcypdnmubavweqr\nndhyrzqaekivtxgcfwosubpm\ncdzatkpnvesrqgomybuwxifh\n\nsujzvcnoy\nzusnxef\ntsqdwghnaulizp\nyusnjbz\nmosjenukz\n\nrqwov\nvqdgorw\novrwq\nopwvqr\n\ndktvrwchflapu\nprlxwcavdfsu\nvadwuqfrpxlc\ncwuvraplfd\n\nqsjpxud\nqpdixsu\n\nxvktceduf\ntuedxcrv\nxcutdev\nedvrxuyct\neqcdxtyuv\n\nsbmhyidelk\neldbhmsiky\nhkmoyisxledb\nkmebisyhdl\nelihmyskdb\n\nycnp\nycnqmsp\n\nrhyjefcnpsqt\nfsqretnhjycp\nefrjhqscptyn\nqscfhrjlypetn\n\njtwuhekpbqaoizgcdym\nomuyzaiedpkwgsqhb\neghavourqipwmdnkbyz\n\nxyejzbknatlrimf\nlrfyqejnxamzdktb\nnatmyzjxevqrfwhkbl\nqmjfzrtlnkeaxby\nycblzpfmektoanjxgr\n\nquwatziscogfnj\nhravz\nxlbmkza\n\nqtjslp\nczikofeuy\nptxbvwal\n\nicmd\nipc\noxftiscgzyb\nciej\nnwkcvqi\n\nmk\no\ndwxb\n\nrzg\nrewg\nrge\nrg\nrwg\n\no\nth\no\n\njtkifz\nltmikzj\nfceztlmipj\npjzati\njgzdbhsytwni\n\nablqpzxmoingerywf\nxpmfezbayrilngoqw\nmixpnyqrfaeglbozw\n\npjewrkfisbo\njkeaprhbfswog\ntojwpfbsekr\nfwbeprkjscmdo\n\nudpjifbsevlwqxgkmtrco\nyimjwpcnukrsdavqxohf\n\nacrljesyth\naontchrvljy\ntalszhrbcfjqey\nchrjtwyqla\nmhxlytjuiagrpkdc\n\nockyq\ndy\n\ne\nex\n\nuz\nzd\nuhz\n\nrqvhze\nhrezqv\nzovfeqhr\n\nboekypjtz\nzktoyjebp\nepbktyojz\nopbktycezj\n\nvpladrzxbwh\npvgizarljdxbwh\n\nsoalndurh\nruhosadbmvf\narudhnos\nlspqurhadoj\n\nualpstjky\npsujtykal\nskaujpylt\ntsapujkyl\n\nhxagklitebnmdrqcfwvsj\nehdxmkqaycvsonbrfwi\n\nxzfmroh\ndntb\ng\n\nvdmbkqil\nsbkwqm\nkqbyrxm\nbkqsm\nmuwkqbzy\n\nqljpfbimaswex\nlwmsbijafpq\nbdalifmpwqsj\njlamqpiwfsb\n\ngtck\nckmwt\n\nwqvp\nws\nwu\nw\nwu\n\nfglbvteoinduycjqksmxzrphwa\nfixyhqwlnkszvoregcudbtpjma\n\noqjb\ngbkqj\njblq\nqbj\n\nyiurdl\nuidlyr\nliuyrd\nruydneli\n\nqznajfvr\nzrjnavfq\nfjqrzanv\n\nndwpjqyx\nyzpdqcnmwxf\n\navnxlomwtjdhqei\ndqfxmreculykgvtjon\n\nfqotuvxki\novywiuqkf\nuiqhkofve\nioukvqf\nquiwvftokp\n\nmcxhanksj\nexnjhocfm\n\nm\nm\ns\nm\nm\n\nmkq\nakmql\nmkq\nkqm\n\nzxldabvrtwoyicp\nrlpbqfwdcovaet\ncrgujdvbkpswmolat\nnwhdbcvpotalr\n\nasngcw\nzxafhgp\nyjegdutvorm\niqngs\nnkgzx\n\nz\nmy\nd\nohfc\ns\n\nk\nr\nm\n\nihelxwpygbo\noebglxywhupi\nlgbxpiyhwo\nbiwoxgplshy\n\nygqnkm\nmknqg\nnkqimgfz\ngnqmk\nglmnqk\n\npeax\nealx\neax\n\nwhgsjnbd\nwhgsnjdb\nwpgdhjnbs\n\npbzmysjve\nebofukvyjpsh\njvyberpwstq\nviyjcdesbwp\nrpsvejbyi\n\nuymlzeihapo\nyleivfaupz\npxzgqlayscbjeiwu\nazeuiyfpl\nryphizkaeul\n\nkyiuectd\nuykdeczi\n\nmf\nzf\nmf\nfn\nf\n\nlksmjwceqfpngyxzd\njmlcsvqdxkyuewngozbp\nwlmeiyqdskgzhjpxnc\njdapwmsxznehqgkcly\n\nxujsqncmrgkflohyewzpa\nqoatelunckgydifmhjpxrvwb\npykmhgazwjncuerloxfq\n\nefyo\nnlcwbjef\n\nnzpwo\nznwcop\n\nwxeguptj\nmgstu\notgu\ngut\nutg\n\ntbpcqyvwizmgs\nscmqvizpytgwb\niqzygwsptvcbm\nybimcptsgqvwz\n\nhlgwbevrqcjxsf\nnjifcqghxrvewtl\ngwjcfmhqexlovrb\n\nxzs\nzxs\n\nfrjzu\nuetgiqm\nuaeov\nsbgyqu\nbquo\n\nyepgsojubnzxrhmtlfq\nhusqezbwxngmrtplf\ngxqneopzrhdslbumft\nhslbzwgeoiuqrnxftpm\nfmunresxbhgaptqzvl\n\nopwuibl\npiluwob\nolwiubp\nwibluop\n\nobvzhrmkn\noprkbvnmzh\nzovnbhmrk\nvnomkzjrbh\nobmkrnzvh\n\nrbgsywpzifanclt\nhfwndizxkjtsuypvcbra\ncznwbamiyptrfs\n\nyopvb\nobpy\nybop\noybp\n\ntcr\ntc\ntc\nct\nct\n\nowztrcxfgy\nftzlorxgwbmk\nwrgoxtfz\n\nqd\nzcn\nmo\n\njpghwqsofrcuatkexy\nsfcxkwaoytjeurhpg\nycwtxehprjfaogsku\nrjfognzhatpecwuykxs\n\nurtsinfxlmqgzhd\nfsrhtuixzgqnmdl\nrtuzixmsfqnhdlg\nyqglzisofmtnrduhx\n\niuawhprfkdqnjsxct\nqihkjosclamtyuvbwnp\nhsdipzjuetkcaqwn\nzpsfhktcqujgixanw\n\ndpey\nzjmyedpx\nnepyd\ndyep\nepnyd\n\nluqyb\nblqtuy\nyjlubq\n\nniulkzgvmpyseqacdjrowb\nlcaozsimjbuqdfygwvker\n\ngzyfw\nzyfgw\nyzgfw\nwyzgf\n\nz\nz\nz\nz\nz\n\nymfsvlrk\nyfmsrvlk\nfmvykslr\nksfrvlym\nvrsklfbym\n\nent\nnwetal\ntne\nnte\n\nnbwkjt\nwbtkj\n\nsydlfim\nyixfumpt\nofkwymbaiz\nqfyendmi\nyfmix\n\nlrdzymgabtexswqn\nxrotnmkzdlgqsc\n\nbilqfktvxanemzgusrph\nsnpakrugbyfmwhqxetizlv\nmehpaftzvlsbqikngxur\n\nkxsnoemdalq\nubptrgfvdcyz\n\nxzvsnur\nsnrzu\n\nejrnailgkmhutwbvdfpyz\ntlmcgefjuvaiyrwkz\n\nftlzhxs\nohxflzts\n\nkzijpbnmfhaxtvl\nflibzjwhtamkxp\nbskirxmthdjflozep\ntujiphbfxmzkl\n\nambtcpgq\nqbmpgtca\nabtqmpgc\nmtpcagqb\n\ntaofbg\ntfoabg\noutgafb\n\nbiurh\nhriqbu\nbiuhr\nrbuzih\nubirh\n\nrvjxdbitgc\ntbrgxjucvid\njmortivxcdbg\ngbrjcvxtdi\n\nqvngfmuyozdbkwaphcr\nhaiufrkzvwydboqjp\nwnukvpfzlhrayoqdb\nzduckowqyebhfsrvpa\ndvrqoswbpaehtyzfuk\n\nyzhmrogndc\nfxzb\nevkwjbpz\njatpvz\nsuqzbi\n\nrxfdvumlhntsowpiycabzgejq\niopsjqmnwhcfuargbytkxvzled\nicaybhzfvpuoeswxqnrlmgdjt\nsneioplcdjfxuvtazqbrgymhw\n\nxjas\nfyhux\n\ngwistylekhruxovq\nlythgqivuowksxer\nigekwlosqturxvy\nroyckiudglevxwsqmt\n\nxvblauri\nrduxqmnt\nxzruo\nuxqrjn\n\nbunsrlgk\nqnyrc\nrwnyq\n\nzkyqfrnbmuatojxvwihel\nzeaxtqjonbrfuklmwivh\ntkbsquaojwvnhzgeriplxmf\nhzikjfvmtruabqexwlon\nlqodamkfxjerunhzibtvw\n\ny\nq\n\nc\ncu\nqkv\n\nervnoqgafw\noscminrewf\nwofjbr\n\nb\nb\nro\nb\nb\n\nylpedbjm\njmyqepfb\npeyjbm\n\nsrba\ngioz\naluw\nbaseu\n\nwamvctoexzgf\nbsyjrxdeo\n\nepbf\nrfobeq\nbfe\netfichnybuv\n\nidyq\nlqiy\nyqil\niyq\nyqi\n\ndhkv\nspcnm\n\nmpuijvlgbd\nbdpfmreatciv\n\nphj\njph\npjh\njpbh\n\nhvox\nodxh\n\ntkiamnuv\n\nenukxohqzridb\ndioqhnrubkxze\nxqokdubnirzhe\niudqxoknhrezb\nhdxrnuqkbseozi\n\nzxjg\nxnak\nkix\n\nqoflemvucnta\nlvaqmncftoeu\nelqmofnvcuat\nyvqlrcafenthom\nvqotfncaelm\n\nczrtsdaiw\nsziwdatcr\n\nvsytnpgefkwxqaouhcmb\noaeldkpvhbcryuxtfsi\nezxktcyfophsabvu\nctuaehvobkyspxf\nyehbrcsovtakpxzuf\n\nwjrd\ndrjw\nrdwj\ndjrw\nwjdr\n\nt\nv\nt\nd\n\nmicqdwj\nrqjmownes\npqtmwj\nyxwjiadmq\n\nwmlsot\nufypiewba\nvtsqzrwdk\n\nxd\nux\n\ngfyrh\nvet\nbml\npka\n\nkonhiwejsydptz\nhtwiogdpnjsy\nsqbfvoywhdjnxtiu\nnstohyjidzwp\n\noulvcfs\ncfvslou\nlvodsucf\ntsvcuolf\nsfoclvu\n\nomfdtxpnqgwhrlac\nqtmxpdngcfarwolh\nptrlngqxchmaodfw\nzxmdcqfrtlhgpaonw\nmtrhldnwpqxgoacf\n\nhfmbnlwtujeqacxo\nbjolymcdfhxaqwu\nmcfahwxjsqboul\nhyobucwmaqlxfj\n\nokxygwhp\nwvkxgpho\nxyufhgwkop\nokgpyhwnbx\n\nct\nc\n\nmrajcxniesvqtzhygfkp\nnmaxyrvsczhipkgejqfot\necmirqgjkhftzpansx\nuaesfjdmixkgqcpthlrzn\nqnpojcfksgtmhzaxrie\n\noinruztqxj\nrvoumlcwdbik\n\nswa\nsayj\nrxygsab\nrzlxs\nusdqnecvmt\n\nkacytqiemxbwrl\nnodjpvuhgt\n\ntu\npte\n\nnyat\nnzft\nrnftd\npotnbg\n\npu\nup\npu\n\nswftrvkn\nfnvwtriks\n\nzvukcgfm\nprdgtfkqh\ngvckxzf\nfkgi\nfkzg\n\nkyojecwftxlq\nkymipznvgh\nwjoukyse\n\njikgxfnpl\nzkupgjrxilmn\n\nl\nl\nral\nl\nl\n\nfeuprzgydcvmniqtbhwsloa\nyzguptfnihbemvqsdjkxc\n\nwlpaenfhz\nehawliznq\n\nnw\nbw\n\nvdtqpkxbrczjfiayumhgo\numoerhdfcazgtj\njldrhtzawcmfugo\nzahcgrtdufjmo\nmuaglcntdjrfzoh\n\nfq\nfq\nvfq\n\nldcsbpxaq\nqcbalszxgd\n\nnzvuexwaqbtkjdhpfls\nvnfdayzbxlqskjuhewp\n\nbejuodklarwyqmftv\nwovrdtyqmbuf\nbfoqtumdywrv\n\npcatgvxsfnemhd\njuyahm\n\nhsgfnxpyirz\npkychzfjasrqni\n\njvhzlgfsreoaqimytdn\nltcaybovrmhinws\n\nhvpaxkrbyf\nyprhbv\nyhzvbpr\nvpjhebymr\nypvbrh\n\nvxbshtzumdg\ngzhdbkvnymt\nmghlcvjibrqewao\n\nqwatcjvglnkoxbz\njlanewzf\newryzpjnla\nhapjrwlnsz\n\nwedijntshvckmuagqroyz\naygdzjuvqetrcnkmhiwos\n\nh\nh\ny\n\ntkzajcrnldp\nptrdkcezajnl\nlrkcdjspnatz\n\nwpbavnk\npbnawkv\nawmygpbfnqkv\nwbvnkap\nanwxvkbp\n\nciojha\nergycubzds\n\nstxylmiduhbek\ncdyluksmixh\nleixuymskhd\n\nckmb\nhcbskxmj\nkcmqba\nckubm\ncbqkm\n\nzimehj\nekhjmz\nzjmehx\nejzpmh\nzadjhemo\n\nuiyfxjtwhzgam\nrzsbqkl\n\nhij\ncxhqiyj\n\ndc\ndi\nnahc\nkzegxs\nhqd\n\ngwcmavtzudlyhbi\ntczmgyaihdvwlu\n\nknmowxtrgysei\nowhxyln\nxnoyw\nowxnvy\n\nawcdeyhxp\nhaidv\najdzh\nroamdnih\ndjhqba\n\nwikxntvrqobmslp\npvskylqtmiouxgerbzwa\n\nvtsodwufbxqagzjcirn\nswbrznlvtdqgcuixoaj\nxnbtowaqukpivrszcdjg\nbdxiqvsoanwtgjcruz\n\ntknlrhyecmdpgavfwxiq\netgzofqjhbxisrlwakm\n\nqlapsncudif\nlcpaubo\nupcla\n\nq\nq\nq\nq\n\nwthxvqcmkrnufij\ncxumyjefhridbtk\nitxcmpfjrukhv\nusqhkciwmxjfzrt\n\novqmbrgjxzehayts\nrvtojeyghszmbaqx\n\nwkdgt\ndtgwk\nztcgwkdf\nkdwtg\n\nbnqdcjklozxsvhaip\ntsboknhvzmpdxa\nbydonswmearvkzpxuh\n\nxouwszgd\nzwuhdo\nouwhnczdxj\nuodkfmqebwz\nuzltdwon\n\nea\nzae\nxae\neax\n\nglpjnaiuqhtcbvxkezsymow\nmyspguazhtokexidjvbfnrw\n\ngtiykcduephanslqxbvzf\nacbtimgdpoefuvlyzxhn\nrltacidghyevpxzbufn\n\nqbnlweychiavojgtpkxm\nzwqlmncaxpouti\nqntcxlismwoarp\n\nywkxe\nsywkex\nyexwk\nkyewx\nwkxye\n\nuiegcdfsh\nugseifdnhc\ndefishcapgumb\n\notaj\nsjdaoqtbc\njatuo\n\nfymlhq\nlifym\nlpkmfhiy\nmuyfl\njlymwcsf\n\nijlhukadenrqsvxgbo\nkcidnlqrhubegxsjavom\nlkjidgnrusewohbvxq\n\nkym\nyk\nyk\n\nyzdjwag\nanxrbdiq\n\nbcmoyupkj\ndhgzlx\n\ntjhagzlncpqe\nhoqbcpzgujnt\n\nrtgifbclnsaehdqouyjk\nwvjdyotzpmhubsaifkclqgre\ntyjasohbcgxfukreqidl\n\njarxzitusmqcd\nwmupfblqejszyd\njdwuqsonmez\n\nlgfksnjtydbomcx\nfmwgnksxbjcldeo\n\nkfiroa\nfrkayo\nnkaflrwo\nrceokmfat\nfrbktyaco\n\nmsrjkzt\ncer\nrd\n\nagvf\nafvg\ngfav\nvfag\n\nmjnwpxdzsg\nlpdjfawbyiz\npuvztdkojwh\ngdzqretwpj\n\nyxnubi\nxibenu\nunbixy\n\ncfhxsvmuebgowktiyqz\nhizkuswbefvcqtmoxlg\nxtyogvkmbwscuizehfrq\ngujihqmopzdfvwxksbtec\n\nvntrdelmwsfhk\nstlgfmrkvdq\nkrmvtdslf\ntksvfrdlm\nefmrkvwdlts\n\nrujwcahnzpe\nbfhzxdpjawtnu\n\nfxjnvwmbgrdeqlh\nrdwgmhevljbxqfn\n\norqvkhczeis\nkieshrpomqvzt\nsvehozikrq\nskrhezvqioc\n\nwxgqml\ngxwslo\nwgysplo\nlhgzw\n\noj\no\no\n\nu\nu\njeu\nu\nu\n\nyiqxngbhocswjrfeudvt\nibygehckxwquomdnjp\nbwcjixudyoqhgeanp\n\nhydfix\nyihxd\nyndxih\nmiwodlxsgehvytk\nnxidhzyb\n\nqaz\nzaquh\nqazu\nqclzba\n\ndvofyzwgjt\nfowjydztgv\nfdzjwgyvto\nyfzdtvwgoj\nvzfyotdgjw\n\na\nswa\na\na\na\n\nlf\nf\nf\nf\n\nehf\nhef\nehf\nhfe\nehf\n\nxpwgek\nujmts\n\nbfzl\nnlkcqzb\n\nclaus\nuacse\nvhucajsk\nlacus\nuasc\n\nbsdpxarynwm\nybawdxpsm\npmbdwasxqy\nxwaqdbpyms\n\nayodzjplqursvtciexbkfm\nxisupbmazlvdkyjterfqc\n\nnqorupyxeicwljsbmhza\nupcrsmlbojzqinxhyaew\npjysnwubqcomaheilzrx\n\nnfilv\nfwbohynv\nelkfvn\nevanf\n\nykzi\nzjmck\nzkm\npzk\nkcz\n\nrcwpgdonmz\nowzdyrcgvnm\nwognrzycmd\nrgblkszmodnciwxq\nwgmornhdzc\n\nmu\num\nrum\num\n\nwdjskfix\nwdkxjfsi\nixwjrfdks\nsdxwjkif\ndsgklwhxifj\n\nwqvfmpctjga\nmqcoafgntj\nqjtgrfcam\nfmawcjgqt\n\nfmqectkvuborwidgj\ngobjtfrqkiumvc\nhnmctagybvzpuriskfqo\n\nbpnwmazxdov\nvnbxadleczwqo\nwxodinbzav\nxwnivzadbmo\ndvoxzanbw\n\npkumcqohndlxstvbijy\nvxwuqmcydjzenpbtlriksf\n\niduobhrnsat\notqruabhdnsz\nditrubnwhs\nkvydthunsrpmcbgl\nbotunsrdh\n\nidyepasroklnwuczhqjg\norijbtsuvlqphfmcgya\n\nzohpdrxsm\ndpjzohfmxr\nprhgocxmzd\nrmhzxlpobdw\nzpocdymgxrh\n\nepjurdmqkiszfow\numpqkgcoeldf\ndmekufqgpo\n\njesvpxdraoyuzctklnqhgi\nfbsqcraxhdzejgilytpvon\neshlqwkraypgojmxvzndcti\ncdlieavgjphtqrusxnozy\ncgotlexyhidaqsvpjmrknz\n\nn\na\nn\n\nimwtv\niuw\nriwx\nixpfruw\n\nghuwqisktymp\nsyuwmgithn\ngwyustmhij\ntrhwyafugsliom\nsphtquwdzygmic\n\nmhrcbwsv\nhmcvrws\nhrwcmvs\n\notm\nbvtmo\nrtosmuxc\n\nfcelpwgamhnquzbsrtdxivjk\ntdjwzsaqhxunkfcvpbrmgil\n"

type PersonGroup struct {
	group []string
}

type AllGroups struct {
	pg []*PersonGroup
}

func cleanString(input string) *AllGroups {
	apg := &AllGroups{}
	splitGroups := strings.SplitAfter(input, "\n\n")

	println(splitGroups)
	for _, j := range splitGroups {

		cleanString := strings.ReplaceAll(j, "\n", " ")

		splitClean := strings.Fields(cleanString)

		p := &PersonGroup{group: splitClean}
		apg.pg = append(apg.pg, p)
	}
	return apg
}

type SingleLetters struct {
	sl []string
}

func (sl *SingleLetters) Insert(s string) *SingleLetters {
	sl.sl = append(sl.sl, s)
	return sl
}

type AllSingleLetter struct {
	ssl        []*SingleLetters
	checkCount int
}

func checkIfAlreadyPresent(s string, ss []string) bool {
	for _, j := range ss {
		if s == j {
			return true
		}
	}
	return false
}

type checkSums struct {
	sums []int
}

func (pg *PersonGroup) checkwhichLettersArePresent() []int {
	sums := checkSums{}

	for _, j := range pg.group {
		var checksum int
		var checkExistent []string
		for i := 0; i < len(j); i++ {
			if !checkIfAlreadyPresent(string([]rune(j)[i]), checkExistent) {

				checkExistent = append(checkExistent, string([]rune(j)[i]))
				checksum++

			} else {

			}

		}
		if checksum == len(j) {
			sums.sums = append(sums.sums, checksum)
		}
	}
	return sums.sums
}

func (pg *PersonGroup) checkwhichLettersArePresent1() []int {
	sums := checkSums{}
	m := map[string]int{}
	var checkExistent []string
	for _, j := range pg.group {
		var checksum int

		for i := 0; i < len(j); i++ {
			if checkIfAlreadyPresent(string([]rune(j)[i]), checkExistent) {
				temp := m[string([]rune(j)[i])]
				temp++
				m[string([]rune(j)[i])] = temp
			} else {
				if !checkIfAlreadyPresent(string([]rune(j)[i]), checkExistent) {

					checkExistent = append(checkExistent, string([]rune(j)[i]))

					m[string([]rune(j)[i])] = 1
				}
			}
		}
		for i, _ := range m {
			if m[i] == len(pg.group) {
				checksum++
			}
		}
		sums.sums = append(sums.sums, checksum)
	}

	return sums.sums
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	temp := []int{}
	apg := cleanString(puzzleInput)
	for _, j := range apg.pg {
		temp2 := sum(j.checkwhichLettersArePresent1())
		temp = append(temp, temp2)

	}
	println(fmt.Sprintf("%v", sum(temp)))

}
