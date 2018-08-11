// +build js,wasm

// Package webgl provides bindings for WebGL via syscall/js.
package webgl

import (
	"errors"

	"syscall/js"
)

func not(v js.Value) bool {
	return v.Type() == js.TypeUndefined || v.Type() == js.TypeNull
}

type ContextAttributes struct {
	Alpha                 bool
	Depth                 bool
	Stencil               bool
	Antialias             bool
	PremultipliedAlpha    bool
	PreserveDrawingBuffer bool
}

func DefaultAttributes() *ContextAttributes {
	return &ContextAttributes{true, true, false, true, true, false}
}

func newContext(v js.Value) *Context {
	ctx := new(Context)
	ctx.Value = v
	ctx.ARRAY_BUFFER = v.Get(`ARRAY_BUFFER`).Int()
	ctx.ARRAY_BUFFER_BINDING = v.Get(`ARRAY_BUFFER_BINDING`).Int()
	ctx.ATTACHED_SHADERS = v.Get(`ATTACHED_SHADERS`).Int()
	ctx.BACK = v.Get(`BACK`).Int()
	ctx.BLEND = v.Get(`BLEND`).Int()
	ctx.BLEND_COLOR = v.Get(`BLEND_COLOR`).Int()
	ctx.BLEND_DST_ALPHA = v.Get(`BLEND_DST_ALPHA`).Int()
	ctx.BLEND_DST_RGB = v.Get(`BLEND_DST_RGB`).Int()
	ctx.BLEND_EQUATION = v.Get(`BLEND_EQUATION`).Int()
	ctx.BLEND_EQUATION_ALPHA = v.Get(`BLEND_EQUATION_ALPHA`).Int()
	ctx.BLEND_EQUATION_RGB = v.Get(`BLEND_EQUATION_RGB`).Int()
	ctx.BLEND_SRC_ALPHA = v.Get(`BLEND_SRC_ALPHA`).Int()
	ctx.BLEND_SRC_RGB = v.Get(`BLEND_SRC_RGB`).Int()
	ctx.BLUE_BITS = v.Get(`BLUE_BITS`).Int()
	ctx.BOOL = v.Get(`BOOL`).Int()
	ctx.BOOL_VEC2 = v.Get(`BOOL_VEC2`).Int()
	ctx.BOOL_VEC3 = v.Get(`BOOL_VEC3`).Int()
	ctx.BOOL_VEC4 = v.Get(`BOOL_VEC4`).Int()
	ctx.BROWSER_DEFAULT_WEBGL = v.Get(`BROWSER_DEFAULT_WEBGL`).Int()
	ctx.BUFFER_SIZE = v.Get(`BUFFER_SIZE`).Int()
	ctx.BUFFER_USAGE = v.Get(`BUFFER_USAGE`).Int()
	ctx.BYTE = v.Get(`BYTE`).Int()
	ctx.CCW = v.Get(`CCW`).Int()
	ctx.CLAMP_TO_EDGE = v.Get(`CLAMP_TO_EDGE`).Int()
	ctx.COLOR_ATTACHMENT0 = v.Get(`COLOR_ATTACHMENT0`).Int()
	ctx.COLOR_BUFFER_BIT = v.Get(`COLOR_BUFFER_BIT`).Int()
	ctx.COLOR_CLEAR_VALUE = v.Get(`COLOR_CLEAR_VALUE`).Int()
	ctx.COLOR_WRITEMASK = v.Get(`COLOR_WRITEMASK`).Int()
	ctx.COMPILE_STATUS = v.Get(`COMPILE_STATUS`).Int()
	ctx.COMPRESSED_TEXTURE_FORMATS = v.Get(`COMPRESSED_TEXTURE_FORMATS`).Int()
	ctx.CONSTANT_ALPHA = v.Get(`CONSTANT_ALPHA`).Int()
	ctx.CONSTANT_COLOR = v.Get(`CONSTANT_COLOR`).Int()
	ctx.CONTEXT_LOST_WEBGL = v.Get(`CONTEXT_LOST_WEBGL`).Int()
	ctx.CULL_FACE = v.Get(`CULL_FACE`).Int()
	ctx.CULL_FACE_MODE = v.Get(`CULL_FACE_MODE`).Int()
	ctx.CURRENT_PROGRAM = v.Get(`CURRENT_PROGRAM`).Int()
	ctx.CURRENT_VERTEX_ATTRIB = v.Get(`CURRENT_VERTEX_ATTRIB`).Int()
	ctx.CW = v.Get(`CW`).Int()
	ctx.DECR = v.Get(`DECR`).Int()
	ctx.DECR_WRAP = v.Get(`DECR_WRAP`).Int()
	ctx.DELETE_STATUS = v.Get(`DELETE_STATUS`).Int()
	ctx.DEPTH_ATTACHMENT = v.Get(`DEPTH_ATTACHMENT`).Int()
	ctx.DEPTH_BITS = v.Get(`DEPTH_BITS`).Int()
	ctx.DEPTH_BUFFER_BIT = v.Get(`DEPTH_BUFFER_BIT`).Int()
	ctx.DEPTH_CLEAR_VALUE = v.Get(`DEPTH_CLEAR_VALUE`).Int()
	ctx.DEPTH_COMPONENT = v.Get(`DEPTH_COMPONENT`).Int()
	ctx.DEPTH_COMPONENT16 = v.Get(`DEPTH_COMPONENT16`).Int()
	ctx.DEPTH_FUNC = v.Get(`DEPTH_FUNC`).Int()
	ctx.DEPTH_RANGE = v.Get(`DEPTH_RANGE`).Int()
	ctx.DEPTH_STENCIL = v.Get(`DEPTH_STENCIL`).Int()
	ctx.DEPTH_STENCIL_ATTACHMENT = v.Get(`DEPTH_STENCIL_ATTACHMENT`).Int()
	ctx.DEPTH_TEST = v.Get(`DEPTH_TEST`).Int()
	ctx.DEPTH_WRITEMASK = v.Get(`DEPTH_WRITEMASK`).Int()
	ctx.DITHER = v.Get(`DITHER`).Int()
	ctx.DONT_CARE = v.Get(`DONT_CARE`).Int()
	ctx.DST_ALPHA = v.Get(`DST_ALPHA`).Int()
	ctx.DST_COLOR = v.Get(`DST_COLOR`).Int()
	ctx.DYNAMIC_DRAW = v.Get(`DYNAMIC_DRAW`).Int()
	ctx.ELEMENT_ARRAY_BUFFER = v.Get(`ELEMENT_ARRAY_BUFFER`).Int()
	ctx.ELEMENT_ARRAY_BUFFER_BINDING = v.Get(`ELEMENT_ARRAY_BUFFER_BINDING`).Int()
	ctx.EQUAL = v.Get(`EQUAL`).Int()
	ctx.FASTEST = v.Get(`FASTEST`).Int()
	ctx.FLOAT = v.Get(`FLOAT`).Int()
	ctx.FLOAT_MAT2 = v.Get(`FLOAT_MAT2`).Int()
	ctx.FLOAT_MAT3 = v.Get(`FLOAT_MAT3`).Int()
	ctx.FLOAT_MAT4 = v.Get(`FLOAT_MAT4`).Int()
	ctx.FLOAT_VEC2 = v.Get(`FLOAT_VEC2`).Int()
	ctx.FLOAT_VEC3 = v.Get(`FLOAT_VEC3`).Int()
	ctx.FLOAT_VEC4 = v.Get(`FLOAT_VEC4`).Int()
	ctx.FRAGMENT_SHADER = v.Get(`FRAGMENT_SHADER`).Int()
	ctx.FRAMEBUFFER = v.Get(`FRAMEBUFFER`).Int()
	ctx.FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = v.Get(`FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE`).Int()
	ctx.FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = v.Get(`FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL`).Int()
	ctx.FRAMEBUFFER_BINDING = v.Get(`FRAMEBUFFER_BINDING`).Int()
	ctx.FRAMEBUFFER_COMPLETE = v.Get(`FRAMEBUFFER_COMPLETE`).Int()
	ctx.FRAMEBUFFER_INCOMPLETE_ATTACHMENT = v.Get(`FRAMEBUFFER_INCOMPLETE_ATTACHMENT`).Int()
	ctx.FRAMEBUFFER_INCOMPLETE_DIMENSIONS = v.Get(`FRAMEBUFFER_INCOMPLETE_DIMENSIONS`).Int()
	ctx.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = v.Get(`FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT`).Int()
	ctx.FRAMEBUFFER_UNSUPPORTED = v.Get(`FRAMEBUFFER_UNSUPPORTED`).Int()
	ctx.FRONT = v.Get(`FRONT`).Int()
	ctx.FRONT_AND_BACK = v.Get(`FRONT_AND_BACK`).Int()
	ctx.FRONT_FACE = v.Get(`FRONT_FACE`).Int()
	ctx.FUNC_ADD = v.Get(`FUNC_ADD`).Int()
	ctx.FUNC_REVERSE_SUBTRACT = v.Get(`FUNC_REVERSE_SUBTRACT`).Int()
	ctx.FUNC_SUBTRACT = v.Get(`FUNC_SUBTRACT`).Int()
	ctx.GENERATE_MIPMAP_HINT = v.Get(`GENERATE_MIPMAP_HINT`).Int()
	ctx.GEQUAL = v.Get(`GEQUAL`).Int()
	ctx.GREATER = v.Get(`GREATER`).Int()
	ctx.GREEN_BITS = v.Get(`GREEN_BITS`).Int()
	ctx.HIGH_FLOAT = v.Get(`HIGH_FLOAT`).Int()
	ctx.HIGH_INT = v.Get(`HIGH_INT`).Int()
	ctx.INCR = v.Get(`INCR`).Int()
	ctx.INCR_WRAP = v.Get(`INCR_WRAP`).Int()
	ctx.INT = v.Get(`INT`).Int()
	ctx.INT_VEC2 = v.Get(`INT_VEC2`).Int()
	ctx.INT_VEC3 = v.Get(`INT_VEC3`).Int()
	ctx.INT_VEC4 = v.Get(`INT_VEC4`).Int()
	ctx.INVALID_ENUM = v.Get(`INVALID_ENUM`).Int()
	ctx.INVALID_FRAMEBUFFER_OPERATION = v.Get(`INVALID_FRAMEBUFFER_OPERATION`).Int()
	ctx.INVALID_OPERATION = v.Get(`INVALID_OPERATION`).Int()
	ctx.INVALID_VALUE = v.Get(`INVALID_VALUE`).Int()
	ctx.INVERT = v.Get(`INVERT`).Int()
	ctx.KEEP = v.Get(`KEEP`).Int()
	ctx.LEQUAL = v.Get(`LEQUAL`).Int()
	ctx.LESS = v.Get(`LESS`).Int()
	ctx.LINEAR = v.Get(`LINEAR`).Int()
	ctx.LINEAR_MIPMAP_LINEAR = v.Get(`LINEAR_MIPMAP_LINEAR`).Int()
	ctx.LINEAR_MIPMAP_NEAREST = v.Get(`LINEAR_MIPMAP_NEAREST`).Int()
	ctx.LINES = v.Get(`LINES`).Int()
	ctx.LINE_LOOP = v.Get(`LINE_LOOP`).Int()
	ctx.LINE_STRIP = v.Get(`LINE_STRIP`).Int()
	ctx.LINE_WIDTH = v.Get(`LINE_WIDTH`).Int()
	ctx.LINK_STATUS = v.Get(`LINK_STATUS`).Int()
	ctx.LOW_FLOAT = v.Get(`LOW_FLOAT`).Int()
	ctx.LOW_INT = v.Get(`LOW_INT`).Int()
	ctx.LUMINANCE = v.Get(`LUMINANCE`).Int()
	ctx.LUMINANCE_ALPHA = v.Get(`LUMINANCE_ALPHA`).Int()
	ctx.MAX_COMBINED_TEXTURE_IMAGE_UNITS = v.Get(`MAX_COMBINED_TEXTURE_IMAGE_UNITS`).Int()
	ctx.MAX_CUBE_MAP_TEXTURE_SIZE = v.Get(`MAX_CUBE_MAP_TEXTURE_SIZE`).Int()
	ctx.MAX_FRAGMENT_UNIFORM_VECTORS = v.Get(`MAX_FRAGMENT_UNIFORM_VECTORS`).Int()
	ctx.MAX_RENDERBUFFER_SIZE = v.Get(`MAX_RENDERBUFFER_SIZE`).Int()
	ctx.MAX_TEXTURE_IMAGE_UNITS = v.Get(`MAX_TEXTURE_IMAGE_UNITS`).Int()
	ctx.MAX_TEXTURE_SIZE = v.Get(`MAX_TEXTURE_SIZE`).Int()
	ctx.MAX_VARYING_VECTORS = v.Get(`MAX_VARYING_VECTORS`).Int()
	ctx.MAX_VERTEX_ATTRIBS = v.Get(`MAX_VERTEX_ATTRIBS`).Int()
	ctx.MAX_VERTEX_TEXTURE_IMAGE_UNITS = v.Get(`MAX_VERTEX_TEXTURE_IMAGE_UNITS`).Int()
	ctx.MAX_VERTEX_UNIFORM_VECTORS = v.Get(`MAX_VERTEX_UNIFORM_VECTORS`).Int()
	ctx.MAX_VIEWPORT_DIMS = v.Get(`MAX_VIEWPORT_DIMS`).Int()
	ctx.MEDIUM_FLOAT = v.Get(`MEDIUM_FLOAT`).Int()
	ctx.MEDIUM_INT = v.Get(`MEDIUM_INT`).Int()
	ctx.MIRRORED_REPEAT = v.Get(`MIRRORED_REPEAT`).Int()
	ctx.NEAREST = v.Get(`NEAREST`).Int()
	ctx.NEAREST_MIPMAP_LINEAR = v.Get(`NEAREST_MIPMAP_LINEAR`).Int()
	ctx.NEAREST_MIPMAP_NEAREST = v.Get(`NEAREST_MIPMAP_NEAREST`).Int()
	ctx.NEVER = v.Get(`NEVER`).Int()
	ctx.NICEST = v.Get(`NICEST`).Int()
	ctx.NONE = v.Get(`NONE`).Int()
	ctx.NOTEQUAL = v.Get(`NOTEQUAL`).Int()
	ctx.NO_ERROR = v.Get(`NO_ERROR`).Int()
	ctx.ONE = v.Get(`ONE`).Int()
	ctx.ONE_MINUS_CONSTANT_ALPHA = v.Get(`ONE_MINUS_CONSTANT_ALPHA`).Int()
	ctx.ONE_MINUS_CONSTANT_COLOR = v.Get(`ONE_MINUS_CONSTANT_COLOR`).Int()
	ctx.ONE_MINUS_DST_ALPHA = v.Get(`ONE_MINUS_DST_ALPHA`).Int()
	ctx.ONE_MINUS_DST_COLOR = v.Get(`ONE_MINUS_DST_COLOR`).Int()
	ctx.ONE_MINUS_SRC_ALPHA = v.Get(`ONE_MINUS_SRC_ALPHA`).Int()
	ctx.ONE_MINUS_SRC_COLOR = v.Get(`ONE_MINUS_SRC_COLOR`).Int()
	ctx.OUT_OF_MEMORY = v.Get(`OUT_OF_MEMORY`).Int()
	ctx.PACK_ALIGNMENT = v.Get(`PACK_ALIGNMENT`).Int()
	ctx.POINTS = v.Get(`POINTS`).Int()
	ctx.POLYGON_OFFSET_FACTOR = v.Get(`POLYGON_OFFSET_FACTOR`).Int()
	ctx.POLYGON_OFFSET_FILL = v.Get(`POLYGON_OFFSET_FILL`).Int()
	ctx.POLYGON_OFFSET_UNITS = v.Get(`POLYGON_OFFSET_UNITS`).Int()
	ctx.RED_BITS = v.Get(`RED_BITS`).Int()
	ctx.RENDERBUFFER = v.Get(`RENDERBUFFER`).Int()
	ctx.RENDERBUFFER_ALPHA_SIZE = v.Get(`RENDERBUFFER_ALPHA_SIZE`).Int()
	ctx.RENDERBUFFER_BINDING = v.Get(`RENDERBUFFER_BINDING`).Int()
	ctx.RENDERBUFFER_BLUE_SIZE = v.Get(`RENDERBUFFER_BLUE_SIZE`).Int()
	ctx.RENDERBUFFER_DEPTH_SIZE = v.Get(`RENDERBUFFER_DEPTH_SIZE`).Int()
	ctx.RENDERBUFFER_GREEN_SIZE = v.Get(`RENDERBUFFER_GREEN_SIZE`).Int()
	ctx.RENDERBUFFER_HEIGHT = v.Get(`RENDERBUFFER_HEIGHT`).Int()
	ctx.RENDERBUFFER_INTERNAL_FORMAT = v.Get(`RENDERBUFFER_INTERNAL_FORMAT`).Int()
	ctx.RENDERBUFFER_RED_SIZE = v.Get(`RENDERBUFFER_RED_SIZE`).Int()
	ctx.RENDERBUFFER_STENCIL_SIZE = v.Get(`RENDERBUFFER_STENCIL_SIZE`).Int()
	ctx.RENDERBUFFER_WIDTH = v.Get(`RENDERBUFFER_WIDTH`).Int()
	ctx.RENDERER = v.Get(`RENDERER`).Int()
	ctx.REPEAT = v.Get(`REPEAT`).Int()
	ctx.REPLACE = v.Get(`REPLACE`).Int()
	ctx.RGB = v.Get(`RGB`).Int()
	ctx.RGB5_A1 = v.Get(`RGB5_A1`).Int()
	ctx.RGB565 = v.Get(`RGB565`).Int()
	ctx.RGBA = v.Get(`RGBA`).Int()
	ctx.RGBA4 = v.Get(`RGBA4`).Int()
	ctx.SAMPLER_2D = v.Get(`SAMPLER_2D`).Int()
	ctx.SAMPLER_CUBE = v.Get(`SAMPLER_CUBE`).Int()
	ctx.SAMPLES = v.Get(`SAMPLES`).Int()
	ctx.SAMPLE_ALPHA_TO_COVERAGE = v.Get(`SAMPLE_ALPHA_TO_COVERAGE`).Int()
	ctx.SAMPLE_BUFFERS = v.Get(`SAMPLE_BUFFERS`).Int()
	ctx.SAMPLE_COVERAGE = v.Get(`SAMPLE_COVERAGE`).Int()
	ctx.SAMPLE_COVERAGE_INVERT = v.Get(`SAMPLE_COVERAGE_INVERT`).Int()
	ctx.SAMPLE_COVERAGE_VALUE = v.Get(`SAMPLE_COVERAGE_VALUE`).Int()
	ctx.SCISSOR_BOX = v.Get(`SCISSOR_BOX`).Int()
	ctx.SCISSOR_TEST = v.Get(`SCISSOR_TEST`).Int()
	ctx.SHADER_TYPE = v.Get(`SHADER_TYPE`).Int()
	ctx.SHADING_LANGUAGE_VERSION = v.Get(`SHADING_LANGUAGE_VERSION`).Int()
	ctx.SHORT = v.Get(`SHORT`).Int()
	ctx.SRC_ALPHA = v.Get(`SRC_ALPHA`).Int()
	ctx.SRC_ALPHA_SATURATE = v.Get(`SRC_ALPHA_SATURATE`).Int()
	ctx.SRC_COLOR = v.Get(`SRC_COLOR`).Int()
	ctx.STATIC_DRAW = v.Get(`STATIC_DRAW`).Int()
	ctx.STENCIL_ATTACHMENT = v.Get(`STENCIL_ATTACHMENT`).Int()
	ctx.STENCIL_BACK_FAIL = v.Get(`STENCIL_BACK_FAIL`).Int()
	ctx.STENCIL_BACK_FUNC = v.Get(`STENCIL_BACK_FUNC`).Int()
	ctx.STENCIL_BACK_PASS_DEPTH_FAIL = v.Get(`STENCIL_BACK_PASS_DEPTH_FAIL`).Int()
	ctx.STENCIL_BACK_PASS_DEPTH_PASS = v.Get(`STENCIL_BACK_PASS_DEPTH_PASS`).Int()
	ctx.STENCIL_BACK_REF = v.Get(`STENCIL_BACK_REF`).Int()
	ctx.STENCIL_BACK_VALUE_MASK = v.Get(`STENCIL_BACK_VALUE_MASK`).Int()
	ctx.STENCIL_BACK_WRITEMASK = v.Get(`STENCIL_BACK_WRITEMASK`).Int()
	ctx.STENCIL_BITS = v.Get(`STENCIL_BITS`).Int()
	ctx.STENCIL_BUFFER_BIT = v.Get(`STENCIL_BUFFER_BIT`).Int()
	ctx.STENCIL_CLEAR_VALUE = v.Get(`STENCIL_CLEAR_VALUE`).Int()
	ctx.STENCIL_FAIL = v.Get(`STENCIL_FAIL`).Int()
	ctx.STENCIL_FUNC = v.Get(`STENCIL_FUNC`).Int()
	ctx.STENCIL_INDEX8 = v.Get(`STENCIL_INDEX8`).Int()
	ctx.STENCIL_PASS_DEPTH_FAIL = v.Get(`STENCIL_PASS_DEPTH_FAIL`).Int()
	ctx.STENCIL_PASS_DEPTH_PASS = v.Get(`STENCIL_PASS_DEPTH_PASS`).Int()
	ctx.STENCIL_REF = v.Get(`STENCIL_REF`).Int()
	ctx.STENCIL_TEST = v.Get(`STENCIL_TEST`).Int()
	ctx.STENCIL_VALUE_MASK = v.Get(`STENCIL_VALUE_MASK`).Int()
	ctx.STENCIL_WRITEMASK = v.Get(`STENCIL_WRITEMASK`).Int()
	ctx.STREAM_DRAW = v.Get(`STREAM_DRAW`).Int()
	ctx.SUBPIXEL_BITS = v.Get(`SUBPIXEL_BITS`).Int()
	ctx.TEXTURE = v.Get(`TEXTURE`).Int()
	ctx.TEXTURE0 = v.Get(`TEXTURE0`).Int()
	ctx.TEXTURE1 = v.Get(`TEXTURE1`).Int()
	ctx.TEXTURE2 = v.Get(`TEXTURE2`).Int()
	ctx.TEXTURE3 = v.Get(`TEXTURE3`).Int()
	ctx.TEXTURE4 = v.Get(`TEXTURE4`).Int()
	ctx.TEXTURE5 = v.Get(`TEXTURE5`).Int()
	ctx.TEXTURE6 = v.Get(`TEXTURE6`).Int()
	ctx.TEXTURE7 = v.Get(`TEXTURE7`).Int()
	ctx.TEXTURE8 = v.Get(`TEXTURE8`).Int()
	ctx.TEXTURE9 = v.Get(`TEXTURE9`).Int()
	ctx.TEXTURE10 = v.Get(`TEXTURE10`).Int()
	ctx.TEXTURE11 = v.Get(`TEXTURE11`).Int()
	ctx.TEXTURE12 = v.Get(`TEXTURE12`).Int()
	ctx.TEXTURE13 = v.Get(`TEXTURE13`).Int()
	ctx.TEXTURE14 = v.Get(`TEXTURE14`).Int()
	ctx.TEXTURE15 = v.Get(`TEXTURE15`).Int()
	ctx.TEXTURE16 = v.Get(`TEXTURE16`).Int()
	ctx.TEXTURE17 = v.Get(`TEXTURE17`).Int()
	ctx.TEXTURE18 = v.Get(`TEXTURE18`).Int()
	ctx.TEXTURE19 = v.Get(`TEXTURE19`).Int()
	ctx.TEXTURE20 = v.Get(`TEXTURE20`).Int()
	ctx.TEXTURE21 = v.Get(`TEXTURE21`).Int()
	ctx.TEXTURE22 = v.Get(`TEXTURE22`).Int()
	ctx.TEXTURE23 = v.Get(`TEXTURE23`).Int()
	ctx.TEXTURE24 = v.Get(`TEXTURE24`).Int()
	ctx.TEXTURE25 = v.Get(`TEXTURE25`).Int()
	ctx.TEXTURE26 = v.Get(`TEXTURE26`).Int()
	ctx.TEXTURE27 = v.Get(`TEXTURE27`).Int()
	ctx.TEXTURE28 = v.Get(`TEXTURE28`).Int()
	ctx.TEXTURE29 = v.Get(`TEXTURE29`).Int()
	ctx.TEXTURE30 = v.Get(`TEXTURE30`).Int()
	ctx.TEXTURE31 = v.Get(`TEXTURE31`).Int()
	ctx.TEXTURE_2D = v.Get(`TEXTURE_2D`).Int()
	ctx.TEXTURE_BINDING_2D = v.Get(`TEXTURE_BINDING_2D`).Int()
	ctx.TEXTURE_BINDING_CUBE_MAP = v.Get(`TEXTURE_BINDING_CUBE_MAP`).Int()
	ctx.TEXTURE_CUBE_MAP = v.Get(`TEXTURE_CUBE_MAP`).Int()
	ctx.TEXTURE_CUBE_MAP_NEGATIVE_X = v.Get(`TEXTURE_CUBE_MAP_NEGATIVE_X`).Int()
	ctx.TEXTURE_CUBE_MAP_NEGATIVE_Y = v.Get(`TEXTURE_CUBE_MAP_NEGATIVE_Y`).Int()
	ctx.TEXTURE_CUBE_MAP_NEGATIVE_Z = v.Get(`TEXTURE_CUBE_MAP_NEGATIVE_Z`).Int()
	ctx.TEXTURE_CUBE_MAP_POSITIVE_X = v.Get(`TEXTURE_CUBE_MAP_POSITIVE_X`).Int()
	ctx.TEXTURE_CUBE_MAP_POSITIVE_Y = v.Get(`TEXTURE_CUBE_MAP_POSITIVE_Y`).Int()
	ctx.TEXTURE_CUBE_MAP_POSITIVE_Z = v.Get(`TEXTURE_CUBE_MAP_POSITIVE_Z`).Int()
	ctx.TEXTURE_MAG_FILTER = v.Get(`TEXTURE_MAG_FILTER`).Int()
	ctx.TEXTURE_MIN_FILTER = v.Get(`TEXTURE_MIN_FILTER`).Int()
	ctx.TEXTURE_WRAP_S = v.Get(`TEXTURE_WRAP_S`).Int()
	ctx.TEXTURE_WRAP_T = v.Get(`TEXTURE_WRAP_T`).Int()
	ctx.TRIANGLES = v.Get(`TRIANGLES`).Int()
	ctx.TRIANGLE_FAN = v.Get(`TRIANGLE_FAN`).Int()
	ctx.TRIANGLE_STRIP = v.Get(`TRIANGLE_STRIP`).Int()
	ctx.UNPACK_ALIGNMENT = v.Get(`UNPACK_ALIGNMENT`).Int()
	ctx.UNPACK_COLORSPACE_CONVERSION_WEBGL = v.Get(`UNPACK_COLORSPACE_CONVERSION_WEBGL`).Int()
	ctx.UNPACK_FLIP_Y_WEBGL = v.Get(`UNPACK_FLIP_Y_WEBGL`).Int()
	ctx.UNPACK_PREMULTIPLY_ALPHA_WEBGL = v.Get(`UNPACK_PREMULTIPLY_ALPHA_WEBGL`).Int()
	ctx.UNSIGNED_BYTE = v.Get(`UNSIGNED_BYTE`).Int()
	ctx.UNSIGNED_INT = v.Get(`UNSIGNED_INT`).Int()
	ctx.UNSIGNED_SHORT = v.Get(`UNSIGNED_SHORT`).Int()
	ctx.UNSIGNED_SHORT_4_4_4_4 = v.Get(`UNSIGNED_SHORT_4_4_4_4`).Int()
	ctx.UNSIGNED_SHORT_5_5_5_1 = v.Get(`UNSIGNED_SHORT_5_5_5_1`).Int()
	ctx.UNSIGNED_SHORT_5_6_5 = v.Get(`UNSIGNED_SHORT_5_6_5`).Int()
	ctx.VALIDATE_STATUS = v.Get(`VALIDATE_STATUS`).Int()
	ctx.VENDOR = v.Get(`VENDOR`).Int()
	ctx.VERSION = v.Get(`VERSION`).Int()
	ctx.VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = v.Get(`VERTEX_ATTRIB_ARRAY_BUFFER_BINDING`).Int()
	ctx.VERTEX_ATTRIB_ARRAY_ENABLED = v.Get(`VERTEX_ATTRIB_ARRAY_ENABLED`).Int()
	ctx.VERTEX_ATTRIB_ARRAY_NORMALIZED = v.Get(`VERTEX_ATTRIB_ARRAY_NORMALIZED`).Int()
	ctx.VERTEX_ATTRIB_ARRAY_POINTER = v.Get(`VERTEX_ATTRIB_ARRAY_POINTER`).Int()
	ctx.VERTEX_ATTRIB_ARRAY_SIZE = v.Get(`VERTEX_ATTRIB_ARRAY_SIZE`).Int()
	ctx.VERTEX_ATTRIB_ARRAY_STRIDE = v.Get(`VERTEX_ATTRIB_ARRAY_STRIDE`).Int()
	ctx.VERTEX_ATTRIB_ARRAY_TYPE = v.Get(`VERTEX_ATTRIB_ARRAY_TYPE`).Int()
	ctx.VERTEX_SHADER = v.Get(`VERTEX_SHADER`).Int()
	ctx.VIEWPORT = v.Get(`VIEWPORT`).Int()
	ctx.ZERO = v.Get(`ZERO`).Int()
	return ctx
}

type Context struct {
	js.Value
	ARRAY_BUFFER                                 int
	ARRAY_BUFFER_BINDING                         int
	ATTACHED_SHADERS                             int
	BACK                                         int
	BLEND                                        int
	BLEND_COLOR                                  int
	BLEND_DST_ALPHA                              int
	BLEND_DST_RGB                                int
	BLEND_EQUATION                               int
	BLEND_EQUATION_ALPHA                         int
	BLEND_EQUATION_RGB                           int
	BLEND_SRC_ALPHA                              int
	BLEND_SRC_RGB                                int
	BLUE_BITS                                    int
	BOOL                                         int
	BOOL_VEC2                                    int
	BOOL_VEC3                                    int
	BOOL_VEC4                                    int
	BROWSER_DEFAULT_WEBGL                        int
	BUFFER_SIZE                                  int
	BUFFER_USAGE                                 int
	BYTE                                         int
	CCW                                          int
	CLAMP_TO_EDGE                                int
	COLOR_ATTACHMENT0                            int
	COLOR_BUFFER_BIT                             int
	COLOR_CLEAR_VALUE                            int
	COLOR_WRITEMASK                              int
	COMPILE_STATUS                               int
	COMPRESSED_TEXTURE_FORMATS                   int
	CONSTANT_ALPHA                               int
	CONSTANT_COLOR                               int
	CONTEXT_LOST_WEBGL                           int
	CULL_FACE                                    int
	CULL_FACE_MODE                               int
	CURRENT_PROGRAM                              int
	CURRENT_VERTEX_ATTRIB                        int
	CW                                           int
	DECR                                         int
	DECR_WRAP                                    int
	DELETE_STATUS                                int
	DEPTH_ATTACHMENT                             int
	DEPTH_BITS                                   int
	DEPTH_BUFFER_BIT                             int
	DEPTH_CLEAR_VALUE                            int
	DEPTH_COMPONENT                              int
	DEPTH_COMPONENT16                            int
	DEPTH_FUNC                                   int
	DEPTH_RANGE                                  int
	DEPTH_STENCIL                                int
	DEPTH_STENCIL_ATTACHMENT                     int
	DEPTH_TEST                                   int
	DEPTH_WRITEMASK                              int
	DITHER                                       int
	DONT_CARE                                    int
	DST_ALPHA                                    int
	DST_COLOR                                    int
	DYNAMIC_DRAW                                 int
	ELEMENT_ARRAY_BUFFER                         int
	ELEMENT_ARRAY_BUFFER_BINDING                 int
	EQUAL                                        int
	FASTEST                                      int
	FLOAT                                        int
	FLOAT_MAT2                                   int
	FLOAT_MAT3                                   int
	FLOAT_MAT4                                   int
	FLOAT_VEC2                                   int
	FLOAT_VEC3                                   int
	FLOAT_VEC4                                   int
	FRAGMENT_SHADER                              int
	FRAMEBUFFER                                  int
	FRAMEBUFFER_ATTACHMENT_Value_NAME            int
	FRAMEBUFFER_ATTACHMENT_Value_TYPE            int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         int
	FRAMEBUFFER_BINDING                          int
	FRAMEBUFFER_COMPLETE                         int
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            int
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS            int
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    int
	FRAMEBUFFER_UNSUPPORTED                      int
	FRONT                                        int
	FRONT_AND_BACK                               int
	FRONT_FACE                                   int
	FUNC_ADD                                     int
	FUNC_REVERSE_SUBTRACT                        int
	FUNC_SUBTRACT                                int
	GENERATE_MIPMAP_HINT                         int
	GEQUAL                                       int
	GREATER                                      int
	GREEN_BITS                                   int
	HIGH_FLOAT                                   int
	HIGH_INT                                     int
	INCR                                         int
	INCR_WRAP                                    int
	INFO_LOG_LENGTH                              int
	INT                                          int
	INT_VEC2                                     int
	INT_VEC3                                     int
	INT_VEC4                                     int
	INVALID_ENUM                                 int
	INVALID_FRAMEBUFFER_OPERATION                int
	INVALID_OPERATION                            int
	INVALID_VALUE                                int
	INVERT                                       int
	KEEP                                         int
	LEQUAL                                       int
	LESS                                         int
	LINEAR                                       int
	LINEAR_MIPMAP_LINEAR                         int
	LINEAR_MIPMAP_NEAREST                        int
	LINES                                        int
	LINE_LOOP                                    int
	LINE_STRIP                                   int
	LINE_WIDTH                                   int
	LINK_STATUS                                  int
	LOW_FLOAT                                    int
	LOW_INT                                      int
	LUMINANCE                                    int
	LUMINANCE_ALPHA                              int
	MAX_COMBINED_TEXTURE_IMAGE_UNITS             int
	MAX_CUBE_MAP_TEXTURE_SIZE                    int
	MAX_FRAGMENT_UNIFORM_VECTORS                 int
	MAX_RENDERBUFFER_SIZE                        int
	MAX_TEXTURE_IMAGE_UNITS                      int
	MAX_TEXTURE_SIZE                             int
	MAX_VARYING_VECTORS                          int
	MAX_VERTEX_ATTRIBS                           int
	MAX_VERTEX_TEXTURE_IMAGE_UNITS               int
	MAX_VERTEX_UNIFORM_VECTORS                   int
	MAX_VIEWPORT_DIMS                            int
	MEDIUM_FLOAT                                 int
	MEDIUM_INT                                   int
	MIRRORED_REPEAT                              int
	NEAREST                                      int
	NEAREST_MIPMAP_LINEAR                        int
	NEAREST_MIPMAP_NEAREST                       int
	NEVER                                        int
	NICEST                                       int
	NONE                                         int
	NOTEQUAL                                     int
	NO_ERROR                                     int
	NUM_COMPRESSED_TEXTURE_FORMATS               int
	ONE                                          int
	ONE_MINUS_CONSTANT_ALPHA                     int
	ONE_MINUS_CONSTANT_COLOR                     int
	ONE_MINUS_DST_ALPHA                          int
	ONE_MINUS_DST_COLOR                          int
	ONE_MINUS_SRC_ALPHA                          int
	ONE_MINUS_SRC_COLOR                          int
	OUT_OF_MEMORY                                int
	PACK_ALIGNMENT                               int
	POINTS                                       int
	POLYGON_OFFSET_FACTOR                        int
	POLYGON_OFFSET_FILL                          int
	POLYGON_OFFSET_UNITS                         int
	RED_BITS                                     int
	RENDERBUFFER                                 int
	RENDERBUFFER_ALPHA_SIZE                      int
	RENDERBUFFER_BINDING                         int
	RENDERBUFFER_BLUE_SIZE                       int
	RENDERBUFFER_DEPTH_SIZE                      int
	RENDERBUFFER_GREEN_SIZE                      int
	RENDERBUFFER_HEIGHT                          int
	RENDERBUFFER_INTERNAL_FORMAT                 int
	RENDERBUFFER_RED_SIZE                        int
	RENDERBUFFER_STENCIL_SIZE                    int
	RENDERBUFFER_WIDTH                           int
	RENDERER                                     int
	REPEAT                                       int
	REPLACE                                      int
	RGB                                          int
	RGB5_A1                                      int
	RGB565                                       int
	RGBA                                         int
	RGBA4                                        int
	SAMPLER_2D                                   int
	SAMPLER_CUBE                                 int
	SAMPLES                                      int
	SAMPLE_ALPHA_TO_COVERAGE                     int
	SAMPLE_BUFFERS                               int
	SAMPLE_COVERAGE                              int
	SAMPLE_COVERAGE_INVERT                       int
	SAMPLE_COVERAGE_VALUE                        int
	SCISSOR_BOX                                  int
	SCISSOR_TEST                                 int
	SHADER_COMPILER                              int
	SHADER_SOURCE_LENGTH                         int
	SHADER_TYPE                                  int
	SHADING_LANGUAGE_VERSION                     int
	SHORT                                        int
	SRC_ALPHA                                    int
	SRC_ALPHA_SATURATE                           int
	SRC_COLOR                                    int
	STATIC_DRAW                                  int
	STENCIL_ATTACHMENT                           int
	STENCIL_BACK_FAIL                            int
	STENCIL_BACK_FUNC                            int
	STENCIL_BACK_PASS_DEPTH_FAIL                 int
	STENCIL_BACK_PASS_DEPTH_PASS                 int
	STENCIL_BACK_REF                             int
	STENCIL_BACK_VALUE_MASK                      int
	STENCIL_BACK_WRITEMASK                       int
	STENCIL_BITS                                 int
	STENCIL_BUFFER_BIT                           int
	STENCIL_CLEAR_VALUE                          int
	STENCIL_FAIL                                 int
	STENCIL_FUNC                                 int
	STENCIL_INDEX                                int
	STENCIL_INDEX8                               int
	STENCIL_PASS_DEPTH_FAIL                      int
	STENCIL_PASS_DEPTH_PASS                      int
	STENCIL_REF                                  int
	STENCIL_TEST                                 int
	STENCIL_VALUE_MASK                           int
	STENCIL_WRITEMASK                            int
	STREAM_DRAW                                  int
	SUBPIXEL_BITS                                int
	TEXTURE                                      int
	TEXTURE0                                     int
	TEXTURE1                                     int
	TEXTURE2                                     int
	TEXTURE3                                     int
	TEXTURE4                                     int
	TEXTURE5                                     int
	TEXTURE6                                     int
	TEXTURE7                                     int
	TEXTURE8                                     int
	TEXTURE9                                     int
	TEXTURE10                                    int
	TEXTURE11                                    int
	TEXTURE12                                    int
	TEXTURE13                                    int
	TEXTURE14                                    int
	TEXTURE15                                    int
	TEXTURE16                                    int
	TEXTURE17                                    int
	TEXTURE18                                    int
	TEXTURE19                                    int
	TEXTURE20                                    int
	TEXTURE21                                    int
	TEXTURE22                                    int
	TEXTURE23                                    int
	TEXTURE24                                    int
	TEXTURE25                                    int
	TEXTURE26                                    int
	TEXTURE27                                    int
	TEXTURE28                                    int
	TEXTURE29                                    int
	TEXTURE30                                    int
	TEXTURE31                                    int
	TEXTURE_2D                                   int
	TEXTURE_BINDING_2D                           int
	TEXTURE_BINDING_CUBE_MAP                     int
	TEXTURE_CUBE_MAP                             int
	TEXTURE_CUBE_MAP_NEGATIVE_X                  int
	TEXTURE_CUBE_MAP_NEGATIVE_Y                  int
	TEXTURE_CUBE_MAP_NEGATIVE_Z                  int
	TEXTURE_CUBE_MAP_POSITIVE_X                  int
	TEXTURE_CUBE_MAP_POSITIVE_Y                  int
	TEXTURE_CUBE_MAP_POSITIVE_Z                  int
	TEXTURE_MAG_FILTER                           int
	TEXTURE_MIN_FILTER                           int
	TEXTURE_WRAP_S                               int
	TEXTURE_WRAP_T                               int
	TRIANGLES                                    int
	TRIANGLE_FAN                                 int
	TRIANGLE_STRIP                               int
	UNPACK_ALIGNMENT                             int
	UNPACK_COLORSPACE_CONVERSION_WEBGL           int
	UNPACK_FLIP_Y_WEBGL                          int
	UNPACK_PREMULTIPLY_ALPHA_WEBGL               int
	UNSIGNED_BYTE                                int
	UNSIGNED_INT                                 int
	UNSIGNED_SHORT                               int
	UNSIGNED_SHORT_4_4_4_4                       int
	UNSIGNED_SHORT_5_5_5_1                       int
	UNSIGNED_SHORT_5_6_5                         int
	VALIDATE_STATUS                              int
	VENDOR                                       int
	VERSION                                      int
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           int
	VERTEX_ATTRIB_ARRAY_ENABLED                  int
	VERTEX_ATTRIB_ARRAY_NORMALIZED               int
	VERTEX_ATTRIB_ARRAY_POINTER                  int
	VERTEX_ATTRIB_ARRAY_SIZE                     int
	VERTEX_ATTRIB_ARRAY_STRIDE                   int
	VERTEX_ATTRIB_ARRAY_TYPE                     int
	VERTEX_SHADER                                int
	VIEWPORT                                     int
	ZERO                                         int
}

func NewContext(canvas js.Value, ca *ContextAttributes) (*Context, error) {
	if not(js.Global().Get("WebGLRenderingContext")) {
		return nil, errors.New("No WebGLRenderingContext")
	}

	if ca == nil {
		ca = DefaultAttributes()
	}

	attrs := map[string]bool{
		"alpha":                 ca.Alpha,
		"depth":                 ca.Depth,
		"stencil":               ca.Stencil,
		"antialias":             ca.Antialias,
		"premultipliedAlpha":    ca.PremultipliedAlpha,
		"preserveDrawingBuffer": ca.PreserveDrawingBuffer,
	}
	atrob := js.Global().Get("Object").New()
	for k, v := range attrs {
		atrob.Set(k, v)
	}
	gl := canvas.Call("getContext", "webgl", atrob)
	if not(gl) {
		return nil, errors.New("Call getContext failed")
	}
	return newContext(gl), nil
}

func (c *Context) GetContextAttributes() ContextAttributes {
	ca := c.Call("getContextAttributes")
	return ContextAttributes{
		ca.Get("alpha").Bool(),
		ca.Get("depth").Bool(),
		ca.Get("stencil").Bool(),
		ca.Get("antialias").Bool(),
		ca.Get("premultipliedAlpha").Bool(),
		ca.Get("preservedDrawingBuffer").Bool(),
	}
}

func (c *Context) ActiveTexture(texture int) {
	c.Call("activeTexture", texture)
}

func (c *Context) AttachShader(program js.Value, shader js.Value) {
	c.Call("attachShader", program, shader)
}

func (c *Context) BindAttribLocation(program js.Value, index int, name string) {
	c.Call("bindAttribLocation", program, index, name)
}

func (c *Context) BindBuffer(target int, buffer js.Value) {
	c.Call("bindBuffer", target, buffer)
}

func (c *Context) BindFramebuffer(target int, framebuffer js.Value) {
	c.Call("bindFramebuffer", target, framebuffer)
}

func (c *Context) BindRenderbuffer(target int, renderbuffer js.Value) {
	c.Call("bindRenderbuffer", target, renderbuffer)
}

func (c *Context) BindTexture(target int, texture js.Value) {
	c.Call("bindTexture", target, texture)
}

func (c *Context) BlendColor(r, g, b, a float64) {
	c.Call("blendColor", r, g, b, a)
}

func (c *Context) BlendEquation(mode int) {
	c.Call("blendEquation", mode)
}

func (c *Context) BlendEquationSeparate(modeRGB, modeAlpha int) {
	c.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

func (c *Context) BlendFunc(sfactor, dfactor int) {
	c.Call("blendFunc", sfactor, dfactor)
}

func (c *Context) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	c.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (c *Context) BufferData(target int, data interface{}, usage int) {
	c.Call("bufferData", target, data, usage)
}

func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	c.Call("bufferSubData", target, offset, data)
}

func (c *Context) CheckFramebufferStatus(target int) int {
	return c.Call("checkFramebufferStatus", target).Int()
}

func (c *Context) Clear(flags int) {
	c.Call("clear", flags)
}

func (c *Context) ClearColor(r, g, b, a float64) {
	c.Call("clearColor", r, g, b, a)
}

func (c *Context) ClearDepth(depth float64) {
	c.Call("clearDepth", depth)
}

func (c *Context) ClearStencil(s int) {
	c.Call("clearStencil", s)
}

func (c *Context) ColorMask(r, g, b, a bool) {
	c.Call("colorMask", r, g, b, a)
}

func (c *Context) CompileShader(shader js.Value) {
	c.Call("compileShader", shader)
}

func (c *Context) CopyTexImage2D(target, level, internal, x, y, w, h, border int) {
	c.Call("copyTexImage2D", target, level, internal, x, y, w, h, border)
}

func (c *Context) CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, w, h int) {
	c.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, w, h)
}

func (c *Context) CreateBuffer() js.Value {
	return c.Call("createBuffer")
}

func (c *Context) CreateFramebuffer() js.Value {
	return c.Call("createFramebuffer")
}

func (c *Context) CreateProgram() js.Value {
	return c.Call("createProgram")
}

func (c *Context) CreateRenderbuffer() js.Value {
	return c.Call("createRenderbuffer")
}

func (c *Context) CreateShader(typ int) js.Value {
	return c.Call("createShader", typ)
}

func (c *Context) CreateTexture() js.Value {
	return c.Call("createTexture")
}

func (c *Context) CullFace(mode int) {
	c.Call("cullFace", mode)
}

func (c *Context) DeleteBuffer(buffer js.Value) {
	c.Call("deleteBuffer", buffer)
}

func (c *Context) DeleteFramebuffer(framebuffer js.Value) {
	c.Call("deleteFramebuffer", framebuffer)
}

func (c *Context) DeleteProgram(program js.Value) {
	c.Call("deleteProgram", program)
}

func (c *Context) DeleteRenderbuffer(renderbuffer js.Value) {
	c.Call("deleteRenderbuffer", renderbuffer)
}

func (c *Context) DeleteShader(shader js.Value) {
	c.Call("deleteShader", shader)
}

func (c *Context) DeleteTexture(texture js.Value) {
	c.Call("deleteTexture", texture)
}

func (c *Context) DepthFunc(fun int) {
	c.Call("depthFunc", fun)
}

func (c *Context) DepthMask(flag bool) {
	c.Call("depthMask", flag)
}

func (c *Context) DepthRange(zNear, zFar float64) {
	c.Call("depthRange", zNear, zFar)
}

func (c *Context) DetachShader(program, shader js.Value) {
	c.Call("detachShader", program, shader)
}

func (c *Context) Disable(cap int) {
	c.Call("disable", cap)
}

func (c *Context) DisableVertexAttribArray(index int) {
	c.Call("disableVertexAttribArray", index)
}

func (c *Context) DrawArrays(mode, first, count int) {
	c.Call("drawArrays", mode, first, count)
}

func (c *Context) DrawElements(mode, count, typ, offset int) {
	c.Call("drawElements", mode, count, typ, offset)
}

func (c *Context) Enable(cap int) {
	c.Call("enable", cap)
}

func (c *Context) EnableVertexAttribArray(index int) {
	c.Call("enableVertexAttribArray", index)
}

func (c *Context) Finish() {
	c.Call("finish")
}

func (c *Context) Flush() {
	c.Call("flush")
}

func (c *Context) FrameBufferRenderBuffer(target, attachment, renderbufferTarget int, renderbuffer js.Value) {
	c.Call("framebufferRenderBuffer", target, attachment, renderbufferTarget, renderbuffer)
}

func (c *Context) FramebufferTexture2D(target, attachment, textarget int, texture js.Value, level int) {
	c.Call("framebufferTexture2D", target, attachment, textarget, texture, level)
}

func (c *Context) FrontFace(mode int) {
	c.Call("frontFace", mode)
}

func (c *Context) GenerateMipmap(target int) {
	c.Call("generateMipmap", target)
}

func (c *Context) GetActiveAttrib(program js.Value, index int) js.Value {
	return c.Call("getActiveAttrib", program, index)
}

func (c *Context) GetActiveUniform(program js.Value, index int) js.Value {
	return c.Call("getActiveUniform", program, index)
}

func (c *Context) GetAttachedShaders(program js.Value) []js.Value {
	objs := c.Call("getAttachedShaders", program)
	shaders := make([]js.Value, objs.Length())
	for i := 0; i < objs.Length(); i++ {
		shaders[i] = objs.Index(i)
	}
	return shaders
}

func (c *Context) GetAttribLocation(program js.Value, name string) int {
	return c.Call("getAttribLocation", program, name).Int()
}

func (c *Context) GetBufferParameter(target, pname int) js.Value {
	return c.Call("getBufferParameter", target, pname)
}

func (c *Context) GetParameter(pname int) js.Value {
	return c.Call("getParameter", pname)
}

func (c *Context) GetError() int {
	return c.Call("getError").Int()
}

func (c *Context) GetExtension(name string) js.Value {
	return c.Call("getExtension", name)
}

func (c *Context) GetFramebufferAttachmentParameter(target, attachment, pname int) js.Value {
	return c.Call("getFramebufferAttachmentParameter", target, attachment, pname)
}

func (c *Context) GetProgramParameteri(program js.Value, pname int) int {
	return c.Call("getProgramParameter", program, pname).Int()
}

func (c *Context) GetProgramParameterb(program js.Value, pname int) bool {
	return c.Call("getProgramParameter", program, pname).Bool()
}

func (c *Context) GetProgramInfoLog(program js.Value) string {
	return c.Call("getProgramInfoLog", program).String()
}

func (c *Context) GetRenderbufferParameter(target, pname int) js.Value {
	return c.Call("getRenderbufferParameter", target, pname)
}

func (c *Context) GetShaderParameter(shader js.Value, pname int) js.Value {
	return c.Call("getShaderParameter", shader, pname)
}

func (c *Context) GetShaderParameterb(shader js.Value, pname int) bool {
	return c.Call("getShaderParameter", shader, pname).Bool()
}

func (c *Context) GetShaderInfoLog(shader js.Value) string {
	return c.Call("getShaderInfoLog", shader).String()
}

func (c *Context) GetShaderSource(shader js.Value) string {
	return c.Call("getShaderSource", shader).String()
}

func (c *Context) GetSupportedExtensions() []string {
	ext := c.Call("getSupportedExtensions")
	extensions := make([]string, ext.Length())
	for i := 0; i < ext.Length(); i++ {
		extensions[i] = ext.Index(i).String()
	}
	return extensions
}

func (c *Context) GetTexParameter(target, pname int) js.Value {
	return c.Call("getTexParameter", target, pname)
}

func (c *Context) GetUniform(program, location js.Value) js.Value {
	return c.Call("getUniform", program, location)
}

func (c *Context) GetUniformLocation(program js.Value, name string) js.Value {
	return c.Call("getUniformLocation", program, name)
}

func (c *Context) GetVertexAttrib(index, pname int) js.Value {
	return c.Call("getVertexAttrib", index, pname)
}

func (c *Context) GetVertexAttribOffset(index, pname int) int {
	return c.Call("getVertexAttribOffset", index, pname).Int()
}

func (c *Context) IsBuffer(buffer js.Value) bool {
	return c.Call("isBuffer", buffer).Bool()
}

func (c *Context) IsContextLost() bool {
	return c.Call("isContextLost").Bool()
}

func (c *Context) IsFramebuffer(framebuffer js.Value) bool {
	return c.Call("isFramebuffer", framebuffer).Bool()
}

func (c *Context) IsProgram(program js.Value) bool {
	return c.Call("isProgram", program).Bool()
}

func (c *Context) IsRenderbuffer(renderbuffer js.Value) bool {
	return c.Call("isRenderbuffer", renderbuffer).Bool()
}

func (c *Context) IsShader(shader js.Value) bool {
	return c.Call("isShader", shader).Bool()
}

func (c *Context) IsTexture(texture js.Value) bool {
	return c.Call("isTexture", texture).Bool()
}

func (c *Context) IsEnabled(capability int) bool {
	return c.Call("isEnabled", capability).Bool()
}

func (c *Context) LineWidth(width float64) {
	c.Call("lineWidth", width)
}

func (c *Context) LinkProgram(program js.Value) {
	c.Call("linkProgram", program)
}

func (c *Context) PixelStorei(pname, param int) {
	c.Call("pixelStorei", pname, param)
}

func (c *Context) PolygonOffset(factor, units float64) {
	c.Call("polygonOffset", factor, units)
}

func (c *Context) ReadPixels(x, y, width, height, format, typ int, pixels js.Value) {
	c.Call("readPixels", x, y, width, height, format, typ, pixels)
}

func (c *Context) RenderbufferStorage(target, internalFormat, width, height int) {
	c.Call("renderbufferStorage", target, internalFormat, width, height)
}

func (c *Context) Scissor(x, y, width, height int) {
	c.Call("scissor", x, y, width, height)
}

func (c *Context) ShaderSource(shader js.Value, source string) {
	c.Call("shaderSource", shader, source)
}

func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, image js.Value) {
	c.Call("texImage2D", target, level, internalFormat, format, kind, image)
}

func (c *Context) TexParameteri(target int, pname int, param int) {
	c.Call("texParameteri", target, pname, param)
}

func (c *Context) TexSubImage2D(target, level, xoffset, yoffset, format, typ int, image js.Value) {
	c.Call("texSubImage2D", target, level, xoffset, yoffset, format, typ, image)
}

func (c *Context) Uniform1f(location js.Value, x float32) {
	c.Call("uniform1f", location, x)
}

func (c *Context) Uniform1i(location js.Value, x int) {
	c.Call("uniform1i", location, x)
}

func (c *Context) Uniform2f(location js.Value, x, y float32) {
	c.Call("uniform2f", location, x, y)
}

func (c *Context) Uniform2i(location js.Value, x, y int) {
	c.Call("uniform2i", location, x, y)
}

func (c *Context) Uniform3f(location js.Value, x, y, z float32) {
	c.Call("uniform3f", location, x, y, z)
}

func (c *Context) Uniform3i(location js.Value, x, y, z int) {
	c.Call("uniform3i", location, x, y, z)
}

func (c *Context) Uniform4f(location js.Value, x, y, z, w float32) {
	c.Call("uniform4f", location, x, y, z, w)
}

func (c *Context) Uniform4i(location js.Value, x, y, z, w int) {
	c.Call("uniform4i", location, x, y, z, w)
}

func (c *Context) UniformMatrix2fv(location js.Value, transpose bool, value []float32) {
	buf := js.TypedArrayOf(value)
	defer buf.Release()
	c.Call("uniformMatrix2fv", location, transpose, buf)
}

func (c *Context) UniformMatrix3fv(location js.Value, transpose bool, value []float32) {
	buf := js.TypedArrayOf(value)
	defer buf.Release()
	c.Call("uniformMatrix3fv", location, transpose, buf)
}

func (c *Context) UniformMatrix4fv(location js.Value, transpose bool, value []float32) {
	buf := js.TypedArrayOf(value)
	defer buf.Release()
	c.Call("uniformMatrix4fv", location, transpose, buf)
}

func (c *Context) UseProgram(program js.Value) {
	c.Call("useProgram", program)
}

func (c *Context) ValidateProgram(program js.Value) {
	c.Call("validateProgram", program)
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	c.Call("vertexAttribPointer", index, size, typ, normal, stride, offset)
}

func (c *Context) Viewport(x, y, width, height int) {
	c.Call("viewport", x, y, width, height)
}
