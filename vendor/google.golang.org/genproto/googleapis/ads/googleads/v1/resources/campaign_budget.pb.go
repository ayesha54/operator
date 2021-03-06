// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_budget.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A campaign budget.
type CampaignBudget struct {
	// The resource name of the campaign budget.
	// Campaign budget resource names have the form:
	//
	// `customers/{customer_id}/campaignBudgets/{budget_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the campaign budget.
	//
	// A campaign budget is created using the CampaignBudgetService create
	// operation and is assigned a budget ID. A budget ID can be shared across
	// different campaigns; the system will then allocate the campaign budget
	// among different campaigns to get optimum results.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the campaign budget.
	//
	// When creating a campaign budget through CampaignBudgetService, every
	// explicitly shared campaign budget must have a non-null, non-empty name.
	// Campaign budgets that are not explicitly shared derive their name from the
	// attached campaign's name.
	//
	// The length of this string must be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The amount of the budget, in the local currency for the account.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit.
	AmountMicros *wrappers.Int64Value `protobuf:"bytes,5,opt,name=amount_micros,json=amountMicros,proto3" json:"amount_micros,omitempty"`
	// The lifetime amount of the budget, in the local currency for the account.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit.
	TotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,10,opt,name=total_amount_micros,json=totalAmountMicros,proto3" json:"total_amount_micros,omitempty"`
	// The status of this campaign budget. This field is read-only.
	Status enums.BudgetStatusEnum_BudgetStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.BudgetStatusEnum_BudgetStatus" json:"status,omitempty"`
	// The delivery method that determines the rate at which the campaign budget
	// is spent.
	//
	// Defaults to STANDARD if unspecified in a create operation.
	DeliveryMethod enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod `protobuf:"varint,7,opt,name=delivery_method,json=deliveryMethod,proto3,enum=google.ads.googleads.v1.enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod" json:"delivery_method,omitempty"`
	// Specifies whether the budget is explicitly shared. Defaults to true if
	// unspecified in a create operation.
	//
	// If true, the budget was created with the purpose of sharing
	// across one or more campaigns.
	//
	// If false, the budget was created with the intention of only being used
	// with a single campaign. The budget's name and status will stay in sync
	// with the campaign's name and status. Attempting to share the budget with a
	// second campaign will result in an error.
	//
	// A non-shared budget can become an explicitly shared. The same operation
	// must also assign the budget a name.
	//
	// A shared campaign budget can never become non-shared.
	ExplicitlyShared *wrappers.BoolValue `protobuf:"bytes,8,opt,name=explicitly_shared,json=explicitlyShared,proto3" json:"explicitly_shared,omitempty"`
	// The number of campaigns actively using the budget.
	//
	// This field is read-only.
	ReferenceCount *wrappers.Int64Value `protobuf:"bytes,9,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	// Indicates whether there is a recommended budget for this campaign budget.
	//
	// This field is read-only.
	HasRecommendedBudget *wrappers.BoolValue `protobuf:"bytes,11,opt,name=has_recommended_budget,json=hasRecommendedBudget,proto3" json:"has_recommended_budget,omitempty"`
	// The recommended budget amount. If no recommendation is available, this will
	// be set to the budget amount.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit.
	//
	// This field is read-only.
	RecommendedBudgetAmountMicros *wrappers.Int64Value `protobuf:"bytes,12,opt,name=recommended_budget_amount_micros,json=recommendedBudgetAmountMicros,proto3" json:"recommended_budget_amount_micros,omitempty"`
	// Period over which to spend the budget. Defaults to DAILY if not specified.
	Period enums.BudgetPeriodEnum_BudgetPeriod `protobuf:"varint,13,opt,name=period,proto3,enum=google.ads.googleads.v1.enums.BudgetPeriodEnum_BudgetPeriod" json:"period,omitempty"`
	// The estimated change in weekly clicks if the recommended budget is applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyClicks *wrappers.Int64Value `protobuf:"bytes,14,opt,name=recommended_budget_estimated_change_weekly_clicks,json=recommendedBudgetEstimatedChangeWeeklyClicks,proto3" json:"recommended_budget_estimated_change_weekly_clicks,omitempty"`
	// The estimated change in weekly cost in micros if the recommended budget is
	// applied. One million is equivalent to one currency unit.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyCostMicros *wrappers.Int64Value `protobuf:"bytes,15,opt,name=recommended_budget_estimated_change_weekly_cost_micros,json=recommendedBudgetEstimatedChangeWeeklyCostMicros,proto3" json:"recommended_budget_estimated_change_weekly_cost_micros,omitempty"`
	// The estimated change in weekly interactions if the recommended budget is
	// applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyInteractions *wrappers.Int64Value `protobuf:"bytes,16,opt,name=recommended_budget_estimated_change_weekly_interactions,json=recommendedBudgetEstimatedChangeWeeklyInteractions,proto3" json:"recommended_budget_estimated_change_weekly_interactions,omitempty"`
	// The estimated change in weekly views if the recommended budget is applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyViews *wrappers.Int64Value `protobuf:"bytes,17,opt,name=recommended_budget_estimated_change_weekly_views,json=recommendedBudgetEstimatedChangeWeeklyViews,proto3" json:"recommended_budget_estimated_change_weekly_views,omitempty"`
	XXX_NoUnkeyedLiteral                        struct{}             `json:"-"`
	XXX_unrecognized                            []byte               `json:"-"`
	XXX_sizecache                               int32                `json:"-"`
}

func (m *CampaignBudget) Reset()         { *m = CampaignBudget{} }
func (m *CampaignBudget) String() string { return proto.CompactTextString(m) }
func (*CampaignBudget) ProtoMessage()    {}
func (*CampaignBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_budget_13a5491b14476a5d, []int{0}
}
func (m *CampaignBudget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBudget.Unmarshal(m, b)
}
func (m *CampaignBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBudget.Marshal(b, m, deterministic)
}
func (dst *CampaignBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBudget.Merge(dst, src)
}
func (m *CampaignBudget) XXX_Size() int {
	return xxx_messageInfo_CampaignBudget.Size(m)
}
func (m *CampaignBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBudget.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBudget proto.InternalMessageInfo

func (m *CampaignBudget) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignBudget) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CampaignBudget) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CampaignBudget) GetAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.AmountMicros
	}
	return nil
}

func (m *CampaignBudget) GetTotalAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TotalAmountMicros
	}
	return nil
}

func (m *CampaignBudget) GetStatus() enums.BudgetStatusEnum_BudgetStatus {
	if m != nil {
		return m.Status
	}
	return enums.BudgetStatusEnum_UNSPECIFIED
}

func (m *CampaignBudget) GetDeliveryMethod() enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod {
	if m != nil {
		return m.DeliveryMethod
	}
	return enums.BudgetDeliveryMethodEnum_UNSPECIFIED
}

func (m *CampaignBudget) GetExplicitlyShared() *wrappers.BoolValue {
	if m != nil {
		return m.ExplicitlyShared
	}
	return nil
}

func (m *CampaignBudget) GetReferenceCount() *wrappers.Int64Value {
	if m != nil {
		return m.ReferenceCount
	}
	return nil
}

func (m *CampaignBudget) GetHasRecommendedBudget() *wrappers.BoolValue {
	if m != nil {
		return m.HasRecommendedBudget
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetAmountMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetAmountMicros
	}
	return nil
}

func (m *CampaignBudget) GetPeriod() enums.BudgetPeriodEnum_BudgetPeriod {
	if m != nil {
		return m.Period
	}
	return enums.BudgetPeriodEnum_UNSPECIFIED
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyClicks() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyClicks
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyCostMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyCostMicros
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyInteractions() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyInteractions
	}
	return nil
}

func (m *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyViews() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedBudgetEstimatedChangeWeeklyViews
	}
	return nil
}

func init() {
	proto.RegisterType((*CampaignBudget)(nil), "google.ads.googleads.v1.resources.CampaignBudget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_budget.proto", fileDescriptor_campaign_budget_13a5491b14476a5d)
}

var fileDescriptor_campaign_budget_13a5491b14476a5d = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdf, 0x4e, 0xdc, 0x46,
	0x14, 0xc6, 0xe5, 0x85, 0xd2, 0x32, 0xc0, 0x02, 0xa6, 0xaa, 0x2c, 0x4a, 0xab, 0xa5, 0x15, 0x12,
	0x12, 0x95, 0xcd, 0xd2, 0x0a, 0x24, 0xb7, 0x17, 0xdd, 0x5d, 0x10, 0xa2, 0x0d, 0xd1, 0x6a, 0x89,
	0x36, 0x52, 0xb4, 0x92, 0x35, 0x78, 0x0e, 0xde, 0x11, 0xf6, 0x8c, 0x35, 0x33, 0x5e, 0xc2, 0x5d,
	0x14, 0xe5, 0x36, 0x52, 0x9e, 0x21, 0x97, 0x79, 0x86, 0x3c, 0x41, 0x1e, 0x25, 0x4f, 0x11, 0xed,
	0xf8, 0x0f, 0xfb, 0x27, 0xc4, 0xbb, 0x77, 0xe3, 0x99, 0xf3, 0xfd, 0xe6, 0x9b, 0xcf, 0xc7, 0x1e,
	0x74, 0x12, 0x70, 0x1e, 0x84, 0xe0, 0x60, 0x22, 0x9d, 0x74, 0x38, 0x1c, 0x0d, 0xea, 0x8e, 0x00,
	0xc9, 0x13, 0xe1, 0x83, 0x74, 0x7c, 0x1c, 0xc5, 0x98, 0x06, 0xcc, 0xbb, 0x4e, 0x48, 0x00, 0xca,
	0x8e, 0x05, 0x57, 0xdc, 0xdc, 0x4d, 0xab, 0x6d, 0x4c, 0xa4, 0x5d, 0x08, 0xed, 0x41, 0xdd, 0x2e,
	0x84, 0xdb, 0xee, 0x63, 0x6c, 0x60, 0x49, 0x24, 0x9d, 0x14, 0xe7, 0x11, 0x08, 0xe9, 0x00, 0xc4,
	0xbd, 0x17, 0x81, 0xea, 0x73, 0x92, 0xe2, 0xb7, 0xeb, 0x33, 0x69, 0x63, 0x10, 0x74, 0x4e, 0x89,
	0x54, 0x58, 0x25, 0x32, 0x93, 0xfc, 0x9a, 0x49, 0xf4, 0xd3, 0x75, 0x72, 0xe3, 0xdc, 0x09, 0x1c,
	0xc7, 0x20, 0xf2, 0xf5, 0x9d, 0x1c, 0x19, 0x53, 0x07, 0x33, 0xc6, 0x15, 0x56, 0x94, 0xb3, 0x6c,
	0xf5, 0xb7, 0x8f, 0x2b, 0xa8, 0xda, 0xca, 0xc2, 0x69, 0x6a, 0xba, 0xf9, 0x3b, 0x5a, 0xcb, 0xcf,
	0xef, 0x31, 0x1c, 0x81, 0x65, 0xd4, 0x8c, 0xfd, 0xe5, 0xce, 0x6a, 0x3e, 0xf9, 0x14, 0x47, 0x60,
	0x1e, 0xa0, 0x0a, 0x25, 0xd6, 0x42, 0xcd, 0xd8, 0x5f, 0x39, 0xfa, 0x39, 0x0b, 0xcf, 0xce, 0x2d,
	0xd8, 0x17, 0x4c, 0x1d, 0xff, 0xd5, 0xc5, 0x61, 0x02, 0x9d, 0x0a, 0x25, 0xe6, 0x21, 0x5a, 0xd4,
	0xa0, 0x45, 0x5d, 0xbe, 0x33, 0x55, 0x7e, 0xa5, 0x04, 0x65, 0x41, 0x5a, 0xaf, 0x2b, 0xcd, 0x7f,
	0xd1, 0x1a, 0x8e, 0x78, 0xc2, 0x94, 0x17, 0x51, 0x5f, 0x70, 0x69, 0x7d, 0x57, 0xbe, 0xd3, 0x6a,
	0xaa, 0xb8, 0xd4, 0x02, 0xf3, 0x7f, 0xb4, 0xa5, 0xb8, 0xc2, 0xa1, 0x37, 0xce, 0x41, 0xe5, 0x9c,
	0x4d, 0xad, 0x6b, 0x8c, 0xc2, 0x9e, 0xa1, 0xa5, 0x34, 0x73, 0x6b, 0xa9, 0x66, 0xec, 0x57, 0x8f,
	0xfe, 0xb1, 0x1f, 0xeb, 0x1c, 0xfd, 0x9e, 0xec, 0x34, 0xc9, 0x2b, 0x2d, 0x39, 0x63, 0x49, 0x34,
	0x36, 0xd1, 0xc9, 0x58, 0x66, 0x82, 0xd6, 0x27, 0x1a, 0xc7, 0xfa, 0x5e, 0xe3, 0x9f, 0xcc, 0x84,
	0x3f, 0xcd, 0xb4, 0x97, 0x5a, 0x3a, 0xb2, 0xcd, 0xf8, 0x42, 0xa7, 0x4a, 0xc6, 0x9e, 0xcd, 0x73,
	0xb4, 0x09, 0x2f, 0xe3, 0x90, 0xfa, 0x54, 0x85, 0xf7, 0x9e, 0xec, 0x63, 0x01, 0xc4, 0xfa, 0x41,
	0xe7, 0xb2, 0x3d, 0x95, 0x4b, 0x93, 0xf3, 0x30, 0x8d, 0x65, 0xe3, 0x41, 0x74, 0xa5, 0x35, 0xe6,
	0x29, 0x5a, 0x17, 0x70, 0x03, 0x02, 0x98, 0x0f, 0x9e, 0x3f, 0x8c, 0xcb, 0x5a, 0x2e, 0x8f, 0xb7,
	0x5a, 0x68, 0x5a, 0x43, 0x89, 0xd9, 0x46, 0x3f, 0xf5, 0xb1, 0xf4, 0x04, 0xf8, 0x3c, 0x8a, 0x80,
	0x11, 0x20, 0xd9, 0x47, 0x6a, 0xad, 0x94, 0x7a, 0xfa, 0xb1, 0x8f, 0x65, 0xe7, 0x41, 0x98, 0x35,
	0x30, 0x41, 0xb5, 0x69, 0xda, 0x44, 0x1f, 0xac, 0x96, 0x1b, 0xfd, 0x45, 0x4c, 0x92, 0x27, 0x7b,
	0x22, 0xfd, 0x74, 0xad, 0xb5, 0x39, 0x7a, 0xa2, 0xad, 0x25, 0x23, 0x2f, 0x2b, 0x9d, 0xe8, 0x64,
	0x2c, 0xf3, 0x8d, 0x81, 0xea, 0x5f, 0x31, 0x0f, 0x52, 0xd1, 0x08, 0x2b, 0x20, 0x9e, 0xdf, 0xc7,
	0x2c, 0x00, 0xef, 0x0e, 0xe0, 0x36, 0xbc, 0xf7, 0xfc, 0x90, 0xfa, 0xb7, 0xd2, 0xaa, 0x96, 0x9f,
	0xe6, 0x8f, 0xa9, 0xd3, 0x9c, 0xe5, 0xcc, 0x96, 0x46, 0x3e, 0xd7, 0xc4, 0x96, 0x06, 0x9a, 0x6f,
	0x0d, 0x74, 0x3c, 0x8f, 0x0d, 0x2e, 0x8b, 0x64, 0xd7, 0xcb, 0xbd, 0x1c, 0xce, 0xe8, 0x85, 0xcb,
	0x3c, 0xec, 0x77, 0x06, 0x3a, 0x99, 0xc3, 0x0f, 0x65, 0x0a, 0x04, 0xf6, 0xf5, 0x8f, 0xce, 0xda,
	0x28, 0x37, 0x74, 0x34, 0x9b, 0xa1, 0x8b, 0x11, 0xac, 0xf9, 0xda, 0x40, 0x87, 0x73, 0x58, 0x1a,
	0x50, 0xb8, 0x93, 0xd6, 0x66, 0xb9, 0x97, 0x83, 0xd9, 0xbc, 0x74, 0x87, 0xbc, 0xe6, 0xab, 0x0a,
	0xda, 0xf3, 0x79, 0x64, 0x97, 0x5e, 0x64, 0xcd, 0xad, 0xf1, 0xbf, 0x7c, 0x7b, 0xb8, 0x73, 0xdb,
	0x78, 0xf1, 0x5f, 0xa6, 0x0c, 0x78, 0x88, 0x59, 0x60, 0x73, 0x11, 0x38, 0x01, 0x30, 0xed, 0x2b,
	0xbf, 0x80, 0x62, 0x2a, 0xbf, 0x71, 0xb5, 0xfe, 0x5d, 0x8c, 0xde, 0x57, 0x16, 0xce, 0x1b, 0x8d,
	0x0f, 0x95, 0xdd, 0xf3, 0x14, 0xd9, 0x20, 0xd2, 0x4e, 0x87, 0xc3, 0x51, 0xb7, 0x6e, 0x77, 0xf2,
	0xca, 0x4f, 0x79, 0x4d, 0xaf, 0x41, 0x64, 0xaf, 0xa8, 0xe9, 0x75, 0xeb, 0xbd, 0xa2, 0xe6, 0x73,
	0x65, 0x2f, 0x5d, 0x70, 0xdd, 0x06, 0x91, 0xae, 0x5b, 0x54, 0xb9, 0x6e, 0xb7, 0xee, 0xba, 0x45,
	0xdd, 0xf5, 0x92, 0x36, 0xfb, 0xe7, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x7f, 0xdd, 0xc4,
	0x06, 0x08, 0x00, 0x00,
}
