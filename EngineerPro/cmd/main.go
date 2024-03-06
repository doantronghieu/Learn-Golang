package cmd

type Config struct {
	DatabasePassword string
}

type UserDataAccessor struct {
	Config Config
}

type OrderDataAccessor struct {
	Config Config
}

type ProductDataAccessor struct {
	Config Config
}

type PostDataAccessor struct {
	Config Config
}

type SocialLogic struct {
	UserDataAccessor UserDataAccessor
	PostDataAccessor PostDataAccessor
}

func NewSocialLogic(
	UserDataAccessor UserDataAccessor,
	PostDataAccessor PostDataAccessor,
) *SocialLogic {
	return &SocialLogic{
		UserDataAccessor: UserDataAccessor,
		PostDataAccessor: PostDataAccessor,
	}
}

type MarketplaceLogic struct {
	UserDataAccessor    UserDataAccessor
	OrderDataAccessor   OrderDataAccessor
	ProductDataAccessor ProductDataAccessor
}

func NewMarketplaceLogic(
	UserDataAccessor UserDataAccessor,
	OrderDataAccessor OrderDataAccessor,
	ProductDataAccessor ProductDataAccessor,
) *MarketplaceLogic {
	return &MarketplaceLogic{
		UserDataAccessor:    UserDataAccessor,
		OrderDataAccessor:   OrderDataAccessor,
		ProductDataAccessor: ProductDataAccessor,
	}
}

type Server struct {
	SocialLogic      SocialLogic
	MarketplaceLogic MarketplaceLogic
}

func NewServer(
	SocialLogic SocialLogic,
	MarketplaceLogic MarketplaceLogic,
) *Server {
	return &Server{
		SocialLogic:      SocialLogic,
		MarketplaceLogic: MarketplaceLogic,
	}
}

func main() {
	// 
}
