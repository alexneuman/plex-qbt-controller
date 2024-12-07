package plex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePlexEventPayload(t *testing.T) {
	parsedJsonBytes := ParsePlexEventDataFromPayload([]byte(payloadString))
	require.NotEmpty(t, parsedJsonBytes)

	plexEventBody := unMarshalPlexEventJSON(parsedJsonBytes)
	assert.NotNil(t, plexEventBody)

}

var payloadString = `b'--------------------------JSBVkHHnqVmnIribq3fN3J\r\nContent-Disposition: form-data; name="payload"\r\nContent-Type: application/json\r\n\r\n{"event":"media.resume","user":false,"owner":true,"Account":{"id":18998100,"thumb":"https://plex.tv/users/c456a86c9c015eba/avatar?c=1733448173","title":"alexshr"},"Server":{"title":"RUBY","uuid":"e99da80b01a4b203868bd66d040d8ccadc5270e6"},"Player":{"local":false,"publicAddress":"172.88.22.70","title":"Alexanders-MacBook-Pro.local","uuid":"sbat2fl5h3m3xdcqscuofqsy"},"Metadata":{"librarySectionType":"movie","ratingKey":"11238","key":"/library/metadata/11238","guid":"plex://movie/5f40cdfabf3e560040d5cd85","slug":"the-super-mario-bros-movie-2023","studio":"Universal Pictures","type":"movie","title":"The Super Mario Bros. Movie","titleSort":"Super Mario Bros. Movie","librarySectionTitle":"RUBY - Movies","librarySectionID":1,"librarySectionKey":"/library/sections/1","contentRating":"PG","summary":"A Brooklyn plumber named Mario travels through the Mushroom Kingdom with a princess named Peach and an anthropomorphic mushroom named Toad to find Mario\'s brother, Luigi, and to save the world from a ruthless fire-breathing Koopa named Bowser.","rating":5.9,"audienceRating":9.5,"viewOffset":3119000,"lastViewedAt":1733448754,"year":2023,"tagline":"Not all heroes wear capes. Some wear overalls.","thumb":"/library/metadata/11238/thumb/1730051708","art":"/library/metadata/11238/art/1730051708","duration":5580000,"originallyAvailableAt":"2023-04-05","addedAt":1684113367,"updatedAt":1730051708,"audienceRatingImage":"rottentomatoes://image.rating.upright","chapterSource":"media","primaryExtraKey":"/library/metadata/11239","ratingImage":"rottentomatoes://image.rating.rotten","Image":[{"alt":"The Super Mario Bros. Movie","type":"coverPoster","url":"/library/metadata/11238/thumb/1730051708"},{"alt":"The Super Mario Bros. Movie","type":"background","url":"/library/metadata/11238/art/1730051708"},{"alt":"The Super Mario Bros. Movie","type":"clearLogo","url":"/library/metadata/11238/clearLogo/1730051708"}],"UltraBlurColors":{"topLeft":"48035c","topRight":"1f336f","bottomRight":"a31e46","bottomLeft":"591718"},"Genre":[{"id":183,"filter":"genre=183","tag":"Adventure","count":365},{"id":310,"filter":"genre=310","tag":"Comedy","count":413},{"id":402,"filter":"genre=402","tag":"Animation","count":161},{"id":879,"filter":"genre=879","tag":"Family","count":185},{"id":403,"filter":"genre=403","tag":"Fantasy","count":259}],"Country":[{"id":429,"filter":"country=429","tag":"Japan","count":46},{"id":53,"filter":"country=53","tag":"United States of America","count":879}],"Guid":[{"id":"imdb://tt6718170"},{"id":"tmdb://502356"},{"id":"tvdb://136578"}],"Rating":[{"image":"imdb://image.rating","value":7.0,"type":"audience","count":974},{"image":"rottentomatoes://image.rating.rotten","value":5.9,"type":"critic","count":213},{"image":"rottentomatoes://image.rating.upright","value":9.5,"type":"audience","count":748},{"image":"themoviedb://image.rating","value":7.7,"type":"audience","count":971}],"Director":[{"id":52090,"filter":"director=52090","tag":"Aaron Horvath","tagKey":"5d776d6d47dd6e001f6f31a8","thumb":"https://metadata-static.plex.tv/5/people/5d8df678154ccfc51234b95b5e8c319e.jpg"},{"id":52089,"filter":"director=52089","tag":"Michael Jelenic","tagKey":"5d77683d2ec6b5001f6bd4db","thumb":"https://metadata-static.plex.tv/3/people/3dcaae0e86639ed9976fc482eaa974ba.jpg"}],"Writer":[{"id":85799,"filter":"writer=85799","tag":"Matthew Fogel","tagKey":"5d77686deb5d26001f1eb0c4","count":2,"thumb":"https://metadata-static.plex.tv/e/people/edf276e442a517df2934a2000a2cf6e0.jpg"}],"Role":[{"id":967,"filter":"actor=967","tag":"Chris Pratt","tagKey":"5d7768328718ba001e313dcc","count":17,"role":"Mario (voice)","thumb":"https://metadata-static.plex.tv/0/people/0d3e21c0f31935b445c1c3ce06503e20.jpg"},{"id":52091,"filter":"actor=52091","tag":"Anya Taylor-Joy","tagKey":"5d776b54ad5437001f79bb69","count":4,"role":"Princess Peach (voice)","thumb":"https://metadata-static.plex.tv/3/people/375bea8abf7bf809f9516be417a5c4c8.jpg"},{"id":35111,"filter":"actor=35111","tag":"Charlie Day","tagKey":"5d77686d3ab0e7001f500cac","count":5,"role":"Luigi (voice)","thumb":"https://metadata-static.plex.tv/6/people/647619e5a2e7fed7865b2d424a8e33f8.jpg"},{"id":24616,"filter":"actor=24616","tag":"Jack Black","tagKey":"5d7768255af944001f1f63d8","count":11,"role":"Bowser (voice)","thumb":"https://metadata-static.plex.tv/f/people/f3f64171d47109215698579164dfb14c.jpg"},{"id":1902,"filter":"actor=1902","tag":"Keegan-Michael Key","tagKey":"5d77683d61141d001fb16771","count":8,"role":"Toad (voice)","thumb":"https://metadata-static.plex.tv/4/people/4fd40e6817141c5193eb289a04755825.jpg"},{"id":740,"filter":"actor=740","tag":"Seth Rogen","tagKey":"5d776825103a2d001f563c1e","count":13,"role":"Donkey Kong (voice)","thumb":"https://metadata-static.plex.tv/b/people/bd52aa5d6d74ac2ffda9a06753ab0631.jpg"},{"id":12858,"filter":"actor=12858","tag":"Fred Armisen","tagKey":"5d77682b8a7581001f12c40e","count":5,"role":"Cranky Kong (voice)","thumb":"https://metadata-static.plex.tv/d/people/d4be2d0400aafb245d1ee2a3020a940a.jpg"},{"id":22943,"filter":"actor=22943","tag":"Sebastian Maniscalco","tagKey":"5d7768612e80df001ebe2a30","count":3,"role":"Spike (voice)","thumb":"https://metadata-static.plex.tv/c/people/ce97caa1511ac9a5f1a3051b8807686c.jpg"},{"id":52092,"filter":"actor=52092","tag":"Charles Martinet","tagKey":"5d77682b103a2d001f5657e7","role":"Mario\'s Dad / Giuseppe (voice)","thumb":"https://metadata-static.plex.tv/d/people/d550d7d4825559ec0679b10d29d5b2a6.jpg"},{"id":1073,"filter":"actor=1073","tag":"Kevin Michael Richardson","tagKey":"5d776829999c64001ec2cf7f","count":13,"role":"Kamek (voice)","thumb":"https://metadata-static.plex.tv/5/people/5fc5962eb25becf7514b127946335cb0.jpg"},{"id":35099,"filter":"actor=35099","tag":"Khary Payton","tagKey":"5d77682a880197001ec9145c","count":2,"role":"Penguin King (voice)","thumb":"https://metadata-static.plex.tv/1/people/148e4914c1b93074c8a40c2689ab9c47.jpg"},{"id":13793,"filter":"actor=13793","tag":"Rino Romano","tagKey":"5d776840999c64001ec31836","role":"Uncle Tony (voice)","thumb":"https://metadata-static.plex.tv/7/people/72dbdd1882caa93e1dbc373f01f18dde.jpg"},{"id":1280,"filter":"actor=1280","tag":"John DiMaggio","tagKey":"5d7768256f4521001ea989ff","count":14,"role":"Uncle Arthur (voice)","thumb":"https://metadata-static.plex.tv/0/people/0a945543418d442fea7ae4948b2a2fda.jpg"},{"id":14444,"filter":"actor=14444","tag":"Jessica DiCicco","tagKey":"5d7768267228e5001f1dce0b","count":4,"role":"Mario\'s Mom (voice)","thumb":"https://metadata-static.plex.tv/d/people/d1128289194faff5500ae55b2f3838d0.jpg"},{"id":1282,"filter":"actor=1282","tag":"Eric Bauza","tagKey":"5d776850e870fd001f19df2b","count":6,"role":"Toad General (voice)","thumb":"https://metadata-static.plex.tv/7/people/71a24cff2a95495a57e4aa5dbe17bb92.jpg"},{"id":85800,"filter":"actor=85800","tag":"Juliet Jelenic","tagKey":"635e315fa38ebd42a0350638","role":"Lumalee (voice)"},{"id":2686,"filter":"actor=2686","tag":"Scott Menville","tagKey":"5d7768337e9a3c0020c6c72a","count":6,"role":"Koopa General (voice)","thumb":"https://metadata-static.plex.tv/5/people/53d22b6b39d9bd485f8426a027133703.jpg"},{"id":1331,"filter":"actor=1331","tag":"Carlos Alazraqui","tagKey":"5d77683485719b001f3a3663","count":9,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/3/people/352966433b141115ce992fe9783483f6.jpg"},{"id":52094,"filter":"actor=52094","tag":"Jason Broad","tagKey":"5f3ff3f53e5306003e57504e","role":"Additional Voices (voice)"},{"id":37850,"filter":"actor=37850","tag":"Ashly Burch","tagKey":"5d7769c8594b2b001e6a9b93","role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/f/people/f456cf63702f07e559630041f9fafe08.jpg"},{"id":52095,"filter":"actor=52095","tag":"Rachel Butera","tagKey":"5d77708e594b2b001e74cfc0","role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/people/5d77708e594b2b001e74cfc0.jpg"},{"id":4523,"filter":"actor=4523","tag":"Cathy Cavadini","tagKey":"5d77682b7e9a3c0020c6b76d","count":6,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/c/people/c863adb54f65dcb90dcef16c2494301c.jpg"},{"id":4518,"filter":"actor=4518","tag":"Will Collyer","tagKey":"5d7768576f4521001eaa2933","count":3,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/f/people/f74f4011f69df240ce72c232e0ccb2d0.jpg"},{"id":52096,"filter":"actor=52096","tag":"Django Craig","tagKey":"5d776834961905001eb93c09","role":"Additional Voices (voice)"},{"id":52097,"filter":"actor=52097","tag":"Willow Geer","tagKey":"5d7768a47a53e9001e6d4ff6","count":2,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/people/5d7768a47a53e9001e6d4ff6.jpg"},{"id":4534,"filter":"actor=4534","tag":"Aaron Hendry","tagKey":"5d77682f6f4521001ea9af78","count":3,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/people/5d77682f6f4521001ea9af78.jpg"},{"id":52098,"filter":"actor=52098","tag":"Andy Hirsch","tagKey":"5d776ae623d5a3001f50a43c","role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/2/people/2e7122ecba4075e74508803effd4ef0c.jpg"},{"id":52104,"filter":"actor=52104","tag":"Barbara Harris","tagKey":"5d77682b7e9a3c0020c6b773","count":2,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/3/people/3e127093c088389f9d6713179a5249ab.jpg"},{"id":1328,"filter":"actor=1328","tag":"Phil LaMarr","tagKey":"5d776827103a2d001f564619","count":11,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/people/5d776827103a2d001f564619.jpg"},{"id":52099,"filter":"actor=52099","tag":"Jeremy Maxwell","tagKey":"5d776855e870fd001f19e251","count":2,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/4/people/4af878f2215a5757b8e74d4d32fd8987.jpg"},{"id":85801,"filter":"actor=85801","tag":"Daniel Mora","tagKey":"5d776b94fb0d55001f56c14f","role":"Additional Voices (voice)"},{"id":130514,"filter":"actor=130514","tag":"Eric E. Osmond","tagKey":"5d776bf5ad5437001f7af645","role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/0/people/0cdd1576e05017612d4e66bf82d9a058.jpg"},{"id":52102,"filter":"actor=52102","tag":"Noreen Reardon","tagKey":"5d776834961905001eb93c13","role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/6/people/6656fa20d3c5739cdf497999c4d7b682.jpg"},{"id":52103,"filter":"actor=52103","tag":"Lee Shorten","tagKey":"5d776b53594b2b001e6d6a66","role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/5/people/5231a7d8fc77c725cc811d103349eacb.jpg"},{"id":888,"filter":"actor=888","tag":"Cree Summer","tagKey":"5d7768337e9a3c0020c6c48c","count":3,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/9/people/9e58cff0835d350e3d279d321794ec75.jpg"},{"id":85802,"filter":"actor=85802","tag":"Nisa Ward","tagKey":"5d77689d6f6af7001ee58ac9","count":2,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/4/people/440e4f45a59ec74edbbdbe4a376c5d89.jpg"},{"id":4520,"filter":"actor=4520","tag":"Nora Wyman","tagKey":"5d776dfd96b655001fe59119","count":2,"role":"Additional Voices (voice)","thumb":"https://metadata-static.plex.tv/4/people/493973da4c9b3a29851da6cca01189c5.jpg"}],"Producer":[{"id":2661,"filter":"producer=2661","tag":"Chris Meledandri","tagKey":"5d7768462e80df001ebe0733","count":7,"thumb":"https://metadata-static.plex.tv/d/people/d0c6aafc9ff38101ab4e0815c8bfb9f0.jpg"},{"id":52105,"filter":"producer=52105","tag":"Shigeru Miyamoto","tagKey":"5d77684654c0f0001f3055a6","thumb":"https://metadata-static.plex.tv/people/5d77684654c0f0001f3055a6.jpg"}],"Label":[{"id":50231,"filter":"label=50231","tag":"Peacock Movies","count":157},{"id":82751,"filter":"label=82751","tag":"Based on a Video Game","count":19},{"id":50225,"filter":"label=50225","tag":"Netflix Movies","count":172}],"Field":[{"locked":true,"name":"thumb"},{"locked":true,"name":"collection"},{"locked":true,"name":"label"}]}}\r\n--------------------------JSBVkHHnqVmnIribq3fN3J--\r\n'`
