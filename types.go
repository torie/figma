package figma

// StrokeAlign specifies where a stroke is drawn relative to the vector outline.
type StrokeAlign string

const (
	StrokeAlignInside  StrokeAlign = "INSIDE"  // draw stroke inside the shape boundary
	StrokeAlignOutside             = "OUTSIDE" // draw stroke outside the shape boundary
	StrokeAlignCenter              = "CENTER"  // draw stroke centered along the shape boundary
)

// LayoutConstraint specifies the constraint relative to the containing Frame.
type LayoutConstraint struct {
	Horizontal HorizontalLayoutConstraint
	Vertical   VerticalLayoutConstraint
}

// HorizontalLayoutConstraint is the layout constraint type for horizontal
// layout.
type HorizontalLayoutConstraint string

const (
	// HorizontalLayoutConstraintTop specifies a node which is laid out relative
	// to top of the containing frame.
	HorizontalLayoutConstraintTop HorizontalLayoutConstraint = "TOP"

	// HorizontalLayoutConstraintBottom specifies a node which is laid out
	// relative to bottom of the containing frame.
	HorizontalLayoutConstraintBottom = "BOTTOM"

	// HorizontalLayoutConstraintCenter specifies a node which is vertically
	// centered relative to containing frame.
	HorizontalLayoutConstraintCenter = "CENTER"

	// HorizontalLayoutConstraintTopBottom specifies a node where the top and
	// bottom of node are constrained relative to containing frame (node
	// stretches with frame).
	HorizontalLayoutConstraintTopBottom = "TOP_BOTTOM"

	// HorizontalLayoutConstraintScale specifies a node which scales vertically
	// with containing frame.
	HorizontalLayoutConstraintScale = "SCALE"
)

// VerticalLayoutConstraint is the layout constraint type for vertical layout.
type VerticalLayoutConstraint string

const (
	// VerticalLayoutConstraintLeft specifies a node which is laid out relative
	// to left of the containing frame.
	VerticalLayoutConstraintLeft VerticalLayoutConstraint = "LEFT"

	// VerticalLayoutConstraintRight specifies a node which is laid out relative
	// to right of the containing frame.
	VerticalLayoutConstraintRight = "RIGHT"

	// VerticalLayoutConstraintCenter specifies a node which is horizontally
	// centered relative to containing frame.
	VerticalLayoutConstraintCenter = "CENTER"

	// VerticalLayoutConstraintLeftRight specifies a node where the left and
	// right side of the node are constrained relative to containing frame (node
	// stretches with frame).
	VerticalLayoutConstraintLeftRight = "LEFT_RIGHT"

	// VerticalLayoutConstraintScale specifies a node which scales horizontally
	// with containing frame.
	VerticalLayoutConstraintScale = "SCALE"
)

// LayoutGrid contains guides to align and place objects within a frame.
type LayoutGrid struct {
	Pattern     LayoutGridPattern
	SectionSize float64
	Visible     bool
	Color       Color
	Alignment   Alignment
	GutterSize  float64
	Offset      float64
	Count       int
}

// Alignment describes positioning of a grid.
type Alignment string

const (
	AlignmentMin    Alignment = "MIN"
	AlignmentMax              = "MAX"
	AlignmentCenter           = "CENTER"
)

// LayoutGridPattern describes the orientation of a grid.
type LayoutGridPattern string

const (
	LayoutGridPatternColumns LayoutGridPattern = "COLUMNS"
	LayoutGridPatternRows                      = "ROWS"
	LayoutGridPatternGrid                      = "GRID"
)

// Effect describes a visual effect such as a shadow or blur.
type Effect struct {
	// Type of effect
	Type EffectType

	// Is the effect active?
	Visible bool

	// Radius of the blur effect (applies to shadows as well)
	Radius float64

	// The following properties are for shadows only:
	// The color of the shadow
	Color Color

	// Blend mode of the shadow
	BlendMode BlendMode

	// How far the shadow is projected in the x and y directions
	Offset Vector
}

// EffectType is the type of effect as a string enum.
type EffectType string

const (
	EffectTypeInnerShadow    EffectType = "INNER_SHADOW"
	EffectTypeDropShadow                = "DROP_SHADOW"
	EffectTypeLayerBlur                 = "LAYER_BLUR"
	EffectTypeBackgroundBlur            = "BACKGROUND_BLUR"
)

// Vector is a 2d vector.
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// ColorStop is a position color pair representing a gradient stop.
type ColorStop struct {
	// Value between 0 and 1 representing position along gradient axis
	Position float64

	// Color attached to corresponding position
	Color Color
}

// BlendMode describes how a layer blends with layers below.
type BlendMode string

const (
	// Normal blends
	BlendModePassThrough = "PASS_THROUGH" // Only applicable to objects with children
	BlendModeNormal      = "NORMAL"

	// Darken
	BlendModeDarken     = "DARKEN"
	BlendModeMultiply   = "MULTIPLY"
	BlendModeLinearBurn = "LINEAR_BURN"
	BlendModeColorBurn  = "COLOR_BURN"

	// Lighten
	BlendModeLighten     = "LIGHTEN"
	BlendModeScreen      = "SCREEN"
	BlendModeLinearDodge = "LINEAR_DODGE"
	BlendModeColorDodge  = "COLOR_DODGE"

	// Contrast
	BlendModeOverlay   = "OVERLAY"
	BlendModeSoftLight = "SOFT_LIGHT"
	BlendModeHardLight = "HARD_LIGHT"

	// Inversion
	BlendModeDifference = "DIFFERENCE"
	BlendModeExclusion  = "EXCLUSION"

	// Component
	BlendModeHue        = "HUE"
	BlendModeSaturation = "SATURATION"
	BlendModeColor      = "COLOR"
	BlendModeLuminosity = "LUMINOSITY"
)

// Rectangle expresses a bounding box in absolute coordinates.
type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// ExportSetting describes the format and size to export an asset at.
type ExportSetting struct {
	Suffix     string
	Format     ImageFormat
	Constraint string
}

// Constraint describes sizing constraint for exports.
type Constraint struct {
	Type  ConstraintType
	Value float64
}

// ConstraintType specifies the type of a constraint.
type ConstraintType string

const (
	ConstraintTypeScale  ConstraintType = "SCALE"
	ConstraintTypeWidth                 = "WIDTH"
	ConstraintTypeHeight                = "HEIGTH"
)

// TypeStyle contains metadata for character formatting.
type TypeStyle struct {
	// Font family of text (standard name)
	FontFamily string `json:"fontFamily"`

	// PostScript font name
	FontPostScriptName string `json:"fontPostScriptName"`

	// Is text italicized?
	Italic bool

	// Numeric font weight
	FontWeight float64 `json:"fontWeight"`

	// Font size in px
	FontSize float64 `json:"fontSize"`

	// Horizontal text alignment
	TextAlignHorizontal HorizontalLayoutConstraint `json:"textAlignHorizontal"`

	// Vertical text alignment
	TextAlignVertical VerticalLayoutConstraint `json:"textAlignVertical"`

	// Space between characters in px
	LetterSpacing float64 `json:"letterSpacing"`

	// Paints applied to characters
	Fills []Paint

	// Line height in px
	LineHeightPx float64 `json:"lineHeightPx"`

	// Line height as a percentage of normal line height
	LineHeightPercent float64 `json:"lineHeightPercent"`
}

// Component is a description of a master component. Helps you identify which
// component instances are attached to.
type Component struct {
	// The name of the component
	Name string

	// The description of the component as entered in the editor
	Description string
}

// Color is an RGBA color.
type Color struct {
	Alpha float64 `json:"a"`
	Red   float64 `json:"r"`
	Green float64 `json:"g"`
	Blue  float64 `json:"b"`
}

// User contains a description of a user.
type User struct {
	//	Name of the user
	Handle string `json:"handle"`

	//	URL link to the user's profile image
	ImgURL string `json:"img_url"`
}

// PaintType specifies the type of paint applied.
type PaintType string

const (
	PaintTypeSolid           PaintType = "SOLID"
	PaintTypeGradientLinear            = "GRADIENT_LINEAR"
	PaintTypeGradientRadial            = "GRADIENT_RADIAL"
	PaintTypeGradientAngular           = "GRADIENT_ANGULAR"
	PaintTypeGradientDiamond           = "GRADIENT_DIAMOND"
	PaintTypeImage                     = "IMAGE"
	PaintTypeEmoji                     = "EMOJI"
)

// ImageFormat specifies the file type of an image.
type ImageFormat string

const (
	ImageFormatPNG ImageFormat = "png"
	ImageFormatSVG             = "svg"
	ImageFormatJPG             = "jpg"
)

// Paint is a solid color, gradient, or image texture that can be applied as
// fills or strokes.
type Paint struct {
	// Type of paint as a string enum
	PaintType PaintType `json:"type"`

	// Is the paint enabled? (default: true)
	Visible bool

	// Overall opacity of paint (colors within the paint can also have opacity
	// values which would blend with this) (default: 1).
	Opacity float64

	// For solid paints:
	// Solid color of the paint
	Color Color

	// For gradient paints: This field contains three vectors, each of which are
	// a position in normalized object space (normalized object space is if the
	// top left corner of the bounding box of the object is (0, 0) and the
	// bottom right is (1,1)). The first position corresponds to the start of
	// the gradient (value 0 for the purposes of calculating gradient stops),
	// the second position is the end of the gradient (value 1), and the third
	// handle position determines the width of the gradient (only relevant for
	// non-linear gradients).
	GradientHandlePositions []Vector

	// Positions of key points along the gradient axis with the colors anchored
	// there. Colors along the gradient are interpolated smoothly between
	// neighboring gradient stops.
	GradientStops []ColorStop

	// For image paints:
	// Image scaling mode
	ScaleMode ScaleMode
}

// ScaleMode specifies the scaling mode of an image.
type ScaleMode string

const (
	ScaleModeFill    ScaleMode = "FILL"
	ScaleModeFit               = "FIT"
	ScaleModeTile              = "TILE"
	ScaleModeStretch           = "STRETCH"
)
