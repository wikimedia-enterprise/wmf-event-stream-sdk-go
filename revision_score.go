package eventstream

import (
	"encoding/json"
	"time"
)

// RevisionScore event scheme struct
type RevisionScore struct {
	baseSchema
	Data struct {
		baseData
		RevParentID  int       `json:"rev_parent_id"`
		RevTimestamp time.Time `json:"rev_timestamp"`
		Scores       struct {
			Articlequality struct {
				ModelName    string   `json:"model_name"`
				ModelVersion string   `json:"model_version"`
				Prediction   []string `json:"prediction"`
				Probability  struct {
					I   float64 `json:"I"`
					II  float64 `json:"II"`
					III float64 `json:"III"`
					IV  float64 `json:"IV"`
				} `json:"probability"`
			} `json:"articlequality"`
			Articletopic struct {
				ModelName    string   `json:"model_name"`
				ModelVersion string   `json:"model_version"`
				Prediction   []string `json:"prediction"`
				Probability  struct {
					CultureBiographyBiography              float64 `json:"Culture.Biography.Biography*"`
					CultureBiographyWomen                  float64 `json:"Culture.Biography.Women"`
					CultureFoodAndDrink                    float64 `json:"Culture.Food and drink"`
					CultureInternetCulture                 float64 `json:"Culture.Internet culture"`
					CultureLinguistics                     float64 `json:"Culture.Linguistics"`
					CultureLiterature                      float64 `json:"Culture.Literature"`
					CultureMediaBooks                      float64 `json:"Culture.Media.Books"`
					CultureMediaEntertainment              float64 `json:"Culture.Media.Entertainment"`
					CultureMediaFilms                      float64 `json:"Culture.Media.Films"`
					CultureMediaMedia                      float64 `json:"Culture.Media.Media*"`
					CultureMediaMusic                      float64 `json:"Culture.Media.Music"`
					CultureMediaRadio                      float64 `json:"Culture.Media.Radio"`
					CultureMediaSoftware                   float64 `json:"Culture.Media.Software"`
					CultureMediaTelevision                 float64 `json:"Culture.Media.Television"`
					CultureMediaVideoGames                 float64 `json:"Culture.Media.Video games"`
					CulturePerformingArts                  float64 `json:"Culture.Performing arts"`
					CulturePhilosophyAndReligion           float64 `json:"Culture.Philosophy and religion"`
					CultureSports                          float64 `json:"Culture.Sports"`
					CultureVisualArtsArchitecture          float64 `json:"Culture.Visual arts.Architecture"`
					CultureVisualArtsComicsAndAnime        float64 `json:"Culture.Visual arts.Comics and Anime"`
					CultureVisualArtsFashion               float64 `json:"Culture.Visual arts.Fashion"`
					CultureVisualArtsVisualArts            float64 `json:"Culture.Visual arts.Visual arts*"`
					GeographyGeographical                  float64 `json:"Geography.Geographical"`
					GeographyRegionsAfricaAfrica           float64 `json:"Geography.Regions.Africa.Africa*"`
					GeographyRegionsAfricaCentralAfrica    float64 `json:"Geography.Regions.Africa.Central Africa"`
					GeographyRegionsAfricaEasternAfrica    float64 `json:"Geography.Regions.Africa.Eastern Africa"`
					GeographyRegionsAfricaNorthernAfrica   float64 `json:"Geography.Regions.Africa.Northern Africa"`
					GeographyRegionsAfricaSouthernAfrica   float64 `json:"Geography.Regions.Africa.Southern Africa"`
					GeographyRegionsAfricaWesternAfrica    float64 `json:"Geography.Regions.Africa.Western Africa"`
					GeographyRegionsAmericasCentralAmerica float64 `json:"Geography.Regions.Americas.Central America"`
					GeographyRegionsAmericasNorthAmerica   float64 `json:"Geography.Regions.Americas.North America"`
					GeographyRegionsAmericasSouthAmerica   float64 `json:"Geography.Regions.Americas.South America"`
					GeographyRegionsAsiaAsia               float64 `json:"Geography.Regions.Asia.Asia*"`
					GeographyRegionsAsiaCentralAsia        float64 `json:"Geography.Regions.Asia.Central Asia"`
					GeographyRegionsAsiaEastAsia           float64 `json:"Geography.Regions.Asia.East Asia"`
					GeographyRegionsAsiaNorthAsia          float64 `json:"Geography.Regions.Asia.North Asia"`
					GeographyRegionsAsiaSouthAsia          float64 `json:"Geography.Regions.Asia.South Asia"`
					GeographyRegionsAsiaSoutheastAsia      float64 `json:"Geography.Regions.Asia.Southeast Asia"`
					GeographyRegionsAsiaWestAsia           float64 `json:"Geography.Regions.Asia.West Asia"`
					GeographyRegionsEuropeEasternEurope    float64 `json:"Geography.Regions.Europe.Eastern Europe"`
					GeographyRegionsEuropeEurope           float64 `json:"Geography.Regions.Europe.Europe*"`
					GeographyRegionsEuropeNorthernEurope   float64 `json:"Geography.Regions.Europe.Northern Europe"`
					GeographyRegionsEuropeSouthernEurope   float64 `json:"Geography.Regions.Europe.Southern Europe"`
					GeographyRegionsEuropeWesternEurope    float64 `json:"Geography.Regions.Europe.Western Europe"`
					GeographyRegionsOceania                float64 `json:"Geography.Regions.Oceania"`
					HistoryAndSocietyBusinessAndEconomics  float64 `json:"History and Society.Business and economics"`
					HistoryAndSocietyEducation             float64 `json:"History and Society.Education"`
					HistoryAndSocietyHistory               float64 `json:"History and Society.History"`
					HistoryAndSocietyMilitaryAndWarfare    float64 `json:"History and Society.Military and warfare"`
					HistoryAndSocietyPoliticsAndGovernment float64 `json:"History and Society.Politics and government"`
					HistoryAndSocietySociety               float64 `json:"History and Society.Society"`
					HistoryAndSocietyTransportation        float64 `json:"History and Society.Transportation"`
					STEMBiology                            float64 `json:"STEM.Biology"`
					STEMChemistry                          float64 `json:"STEM.Chemistry"`
					STEMComputing                          float64 `json:"STEM.Computing"`
					STEMEarthAndEnvironment                float64 `json:"STEM.Earth and environment"`
					STEMEngineering                        float64 `json:"STEM.Engineering"`
					STEMLibrariesInformation               float64 `json:"STEM.Libraries & Information"`
					STEMMathematics                        float64 `json:"STEM.Mathematics"`
					STEMMedicineHealth                     float64 `json:"STEM.Medicine & Health"`
					STEMPhysics                            float64 `json:"STEM.Physics"`
					STEMSTEM                               float64 `json:"STEM.STEM*"`
					STEMSpace                              float64 `json:"STEM.Space"`
					STEMTechnology                         float64 `json:"STEM.Technology"`
				} `json:"probability"`
			} `json:"articletopic"`
			Damaging struct {
				ModelName    string   `json:"model_name"`
				ModelVersion string   `json:"model_version"`
				Prediction   []string `json:"prediction"`
				Probability  struct {
					False float64 `json:"false"`
					True  float64 `json:"true"`
				} `json:"probability"`
			} `json:"damaging"`
			Drafttopic struct {
				ModelName    string   `json:"model_name"`
				ModelVersion string   `json:"model_version"`
				Prediction   []string `json:"prediction"`
				Probability  struct {
					CultureBiographyBiography              float64 `json:"Culture.Biography.Biography*"`
					CultureBiographyWomen                  float64 `json:"Culture.Biography.Women"`
					CultureFoodAndDrink                    float64 `json:"Culture.Food and drink"`
					CultureInternetCulture                 float64 `json:"Culture.Internet culture"`
					CultureLinguistics                     float64 `json:"Culture.Linguistics"`
					CultureLiterature                      float64 `json:"Culture.Literature"`
					CultureMediaBooks                      float64 `json:"Culture.Media.Books"`
					CultureMediaEntertainment              float64 `json:"Culture.Media.Entertainment"`
					CultureMediaFilms                      float64 `json:"Culture.Media.Films"`
					CultureMediaMedia                      float64 `json:"Culture.Media.Media*"`
					CultureMediaMusic                      float64 `json:"Culture.Media.Music"`
					CultureMediaRadio                      float64 `json:"Culture.Media.Radio"`
					CultureMediaSoftware                   float64 `json:"Culture.Media.Software"`
					CultureMediaTelevision                 float64 `json:"Culture.Media.Television"`
					CultureMediaVideoGames                 float64 `json:"Culture.Media.Video games"`
					CulturePerformingArts                  float64 `json:"Culture.Performing arts"`
					CulturePhilosophyAndReligion           float64 `json:"Culture.Philosophy and religion"`
					CultureSports                          float64 `json:"Culture.Sports"`
					CultureVisualArtsArchitecture          float64 `json:"Culture.Visual arts.Architecture"`
					CultureVisualArtsComicsAndAnime        float64 `json:"Culture.Visual arts.Comics and Anime"`
					CultureVisualArtsFashion               float64 `json:"Culture.Visual arts.Fashion"`
					CultureVisualArtsVisualArts            float64 `json:"Culture.Visual arts.Visual arts*"`
					GeographyGeographical                  float64 `json:"Geography.Geographical"`
					GeographyRegionsAfricaAfrica           float64 `json:"Geography.Regions.Africa.Africa*"`
					GeographyRegionsAfricaCentralAfrica    float64 `json:"Geography.Regions.Africa.Central Africa"`
					GeographyRegionsAfricaEasternAfrica    float64 `json:"Geography.Regions.Africa.Eastern Africa"`
					GeographyRegionsAfricaNorthernAfrica   float64 `json:"Geography.Regions.Africa.Northern Africa"`
					GeographyRegionsAfricaSouthernAfrica   float64 `json:"Geography.Regions.Africa.Southern Africa"`
					GeographyRegionsAfricaWesternAfrica    float64 `json:"Geography.Regions.Africa.Western Africa"`
					GeographyRegionsAmericasCentralAmerica float64 `json:"Geography.Regions.Americas.Central America"`
					GeographyRegionsAmericasNorthAmerica   float64 `json:"Geography.Regions.Americas.North America"`
					GeographyRegionsAmericasSouthAmerica   float64 `json:"Geography.Regions.Americas.South America"`
					GeographyRegionsAsiaAsia               float64 `json:"Geography.Regions.Asia.Asia*"`
					GeographyRegionsAsiaCentralAsia        float64 `json:"Geography.Regions.Asia.Central Asia"`
					GeographyRegionsAsiaEastAsia           float64 `json:"Geography.Regions.Asia.East Asia"`
					GeographyRegionsAsiaNorthAsia          float64 `json:"Geography.Regions.Asia.North Asia"`
					GeographyRegionsAsiaSouthAsia          float64 `json:"Geography.Regions.Asia.South Asia"`
					GeographyRegionsAsiaSoutheastAsia      float64 `json:"Geography.Regions.Asia.Southeast Asia"`
					GeographyRegionsAsiaWestAsia           float64 `json:"Geography.Regions.Asia.West Asia"`
					GeographyRegionsEuropeEasternEurope    float64 `json:"Geography.Regions.Europe.Eastern Europe"`
					GeographyRegionsEuropeEurope           float64 `json:"Geography.Regions.Europe.Europe*"`
					GeographyRegionsEuropeNorthernEurope   float64 `json:"Geography.Regions.Europe.Northern Europe"`
					GeographyRegionsEuropeSouthernEurope   float64 `json:"Geography.Regions.Europe.Southern Europe"`
					GeographyRegionsEuropeWesternEurope    float64 `json:"Geography.Regions.Europe.Western Europe"`
					GeographyRegionsOceania                float64 `json:"Geography.Regions.Oceania"`
					HistoryAndSocietyBusinessAndEconomics  float64 `json:"History and Society.Business and economics"`
					HistoryAndSocietyEducation             float64 `json:"History and Society.Education"`
					HistoryAndSocietyHistory               float64 `json:"History and Society.History"`
					HistoryAndSocietyMilitaryAndWarfare    float64 `json:"History and Society.Military and warfare"`
					HistoryAndSocietyPoliticsAndGovernment float64 `json:"History and Society.Politics and government"`
					HistoryAndSocietySociety               float64 `json:"History and Society.Society"`
					HistoryAndSocietyTransportation        float64 `json:"History and Society.Transportation"`
					STEMBiology                            float64 `json:"STEM.Biology"`
					STEMChemistry                          float64 `json:"STEM.Chemistry"`
					STEMComputing                          float64 `json:"STEM.Computing"`
					STEMEarthAndEnvironment                float64 `json:"STEM.Earth and environment"`
					STEMEngineering                        float64 `json:"STEM.Engineering"`
					STEMLibrariesInformation               float64 `json:"STEM.Libraries & Information"`
					STEMMathematics                        float64 `json:"STEM.Mathematics"`
					STEMMedicineHealth                     float64 `json:"STEM.Medicine & Health"`
					STEMPhysics                            float64 `json:"STEM.Physics"`
					STEMSTEM                               float64 `json:"STEM.STEM*"`
					STEMSpace                              float64 `json:"STEM.Space"`
					STEMTechnology                         float64 `json:"STEM.Technology"`
				} `json:"probability"`
			} `json:"drafttopic"`
			Goodfaith struct {
				ModelName    string   `json:"model_name"`
				ModelVersion string   `json:"model_version"`
				Prediction   []string `json:"prediction"`
				Probability  struct {
					False float64 `json:"false"`
					True  float64 `json:"true"`
				} `json:"probability"`
			} `json:"goodfaith"`
			Itemquality struct {
				ModelName    string   `json:"model_name"`
				ModelVersion string   `json:"model_version"`
				Prediction   []string `json:"prediction"`
				Probability  struct {
					A float64 `json:"A"`
					B float64 `json:"B"`
					C float64 `json:"C"`
					D float64 `json:"D"`
					E float64 `json:"E"`
				} `json:"probability"`
			} `json:"itemquality"`
			Reverted struct {
				ModelName    string   `json:"model_name"`
				ModelVersion string   `json:"model_version"`
				Prediction   []string `json:"prediction"`
				Probability  struct {
					False float64 `json:"false"`
					True  float64 `json:"true"`
				} `json:"probability"`
			} `json:"reverted"`
		} `json:"scores"`
	}
}

func (rs *RevisionScore) timestamp() time.Time {
	return rs.Data.Meta.Dt
}

func (rs *RevisionScore) unmarshal(evt *Event) error {
	rs.ID = evt.ID
	return json.Unmarshal(evt.Data, &rs.Data)
}
