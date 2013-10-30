package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var adddmi_code cu.Function

type adddmi_args struct {
	arg_Hx unsafe.Pointer
	arg_Hy unsafe.Pointer
	arg_Hz unsafe.Pointer
	arg_mx unsafe.Pointer
	arg_my unsafe.Pointer
	arg_mz unsafe.Pointer
	arg_Dx float32
	arg_Dy float32
	arg_Dz float32
	arg_A  float32
	arg_cx float32
	arg_cy float32
	arg_cz float32
	arg_Nx int
	arg_Ny int
	arg_Nz int
	argptr [16]unsafe.Pointer
}

// Wrapper for adddmi CUDA kernel, asynchronous.
func k_adddmi_async(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, A float32, cx float32, cy float32, cz float32, Nx int, Ny int, Nz int, cfg *config, str int) {
	if adddmi_code == 0 {
		adddmi_code = fatbinLoad(adddmi_map, "adddmi")
	}

	var _a_ adddmi_args

	_a_.arg_Hx = Hx
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_Hx)
	_a_.arg_Hy = Hy
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_Hy)
	_a_.arg_Hz = Hz
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_Hz)
	_a_.arg_mx = mx
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_mx)
	_a_.arg_my = my
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_my)
	_a_.arg_mz = mz
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_mz)
	_a_.arg_Dx = Dx
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_Dx)
	_a_.arg_Dy = Dy
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_Dy)
	_a_.arg_Dz = Dz
	_a_.argptr[8] = unsafe.Pointer(&_a_.arg_Dz)
	_a_.arg_A = A
	_a_.argptr[9] = unsafe.Pointer(&_a_.arg_A)
	_a_.arg_cx = cx
	_a_.argptr[10] = unsafe.Pointer(&_a_.arg_cx)
	_a_.arg_cy = cy
	_a_.argptr[11] = unsafe.Pointer(&_a_.arg_cy)
	_a_.arg_cz = cz
	_a_.argptr[12] = unsafe.Pointer(&_a_.arg_cz)
	_a_.arg_Nx = Nx
	_a_.argptr[13] = unsafe.Pointer(&_a_.arg_Nx)
	_a_.arg_Ny = Ny
	_a_.argptr[14] = unsafe.Pointer(&_a_.arg_Ny)
	_a_.arg_Nz = Nz
	_a_.argptr[15] = unsafe.Pointer(&_a_.arg_Nz)

	args := _a_.argptr[:]
	cu.LaunchKernel(adddmi_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)
}

// Wrapper for adddmi CUDA kernel, synchronized.
func k_adddmi(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, A float32, cx float32, cy float32, cz float32, Nx int, Ny int, Nz int, cfg *config) {
	const stream = 0
	k_adddmi_async(Hx, Hy, Hz, mx, my, mz, Dx, Dy, Dz, A, cx, cy, cz, Nx, Ny, Nz, cfg, stream)
	Sync(stream)
}

var adddmi_map = map[int]string{0: "",
	20: adddmi_ptx_20,
	30: adddmi_ptx_30,
	35: adddmi_ptx_35}

const (
	adddmi_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .f32 adddmi_param_9,
	.param .f32 adddmi_param_10,
	.param .f32 adddmi_param_11,
	.param .f32 adddmi_param_12,
	.param .u32 adddmi_param_13,
	.param .u32 adddmi_param_14,
	.param .u32 adddmi_param_15
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<104>;
	.reg .f32 	%f<149>;
	.reg .s64 	%rd<55>;


	ld.param.u64 	%rd7, [adddmi_param_0];
	ld.param.u64 	%rd8, [adddmi_param_1];
	ld.param.u64 	%rd9, [adddmi_param_2];
	ld.param.u64 	%rd10, [adddmi_param_3];
	ld.param.u64 	%rd11, [adddmi_param_4];
	ld.param.u64 	%rd12, [adddmi_param_5];
	ld.param.f32 	%f35, [adddmi_param_6];
	ld.param.f32 	%f36, [adddmi_param_7];
	ld.param.f32 	%f37, [adddmi_param_9];
	ld.param.f32 	%f38, [adddmi_param_10];
	ld.param.f32 	%f39, [adddmi_param_11];
	ld.param.u32 	%r4, [adddmi_param_13];
	ld.param.u32 	%r5, [adddmi_param_14];
	ld.param.u32 	%r6, [adddmi_param_15];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd9;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd7;
	.loc 1 17 1
	mov.u32 	%r7, %ntid.z;
	mov.u32 	%r8, %ctaid.z;
	mov.u32 	%r9, %tid.z;
	mad.lo.s32 	%r1, %r7, %r8, %r9;
	.loc 1 18 1
	mov.u32 	%r10, %ntid.y;
	mov.u32 	%r11, %ctaid.y;
	mov.u32 	%r12, %tid.y;
	mad.lo.s32 	%r2, %r10, %r11, %r12;
	.loc 1 19 1
	mov.u32 	%r13, %ntid.x;
	mov.u32 	%r14, %ctaid.x;
	mov.u32 	%r15, %tid.x;
	mad.lo.s32 	%r3, %r13, %r14, %r15;
	.loc 1 21 1
	setp.ge.s32	%p1, %r2, %r5;
	setp.ge.s32	%p2, %r3, %r4;
	or.pred  	%p3, %p2, %p1;
	.loc 1 21 1
	setp.ge.s32	%p4, %r1, %r6;
	or.pred  	%p5, %p3, %p4;
	.loc 1 21 1
	@%p5 bra 	BB0_14;

	.loc 1 25 1
	mad.lo.s32 	%r16, %r1, %r5, %r2;
	mad.lo.s32 	%r17, %r16, %r4, %r3;
	mul.wide.s32 	%rd13, %r17, 4;
	add.s64 	%rd14, %rd6, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f1, [%rd14];
	add.s64 	%rd15, %rd5, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f2, [%rd15];
	add.s64 	%rd16, %rd4, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd3, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f4, [%rd17];
	add.s64 	%rd18, %rd2, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f5, [%rd18];
	add.s64 	%rd19, %rd1, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f6, [%rd19];
	.loc 1 36 1
	add.s32 	%r18, %r3, 1;
	setp.lt.s32	%p6, %r18, %r4;
	@%p6 bra 	BB0_3;

	mov.f32 	%f139, 0f00000000;
	mov.f32 	%f138, %f139;
	mov.f32 	%f137, %f139;
	bra.uni 	BB0_4;

BB0_3:
	mul.wide.s32 	%rd21, %r17, 4;
	add.s64 	%rd22, %rd3, %rd21;
	.loc 1 38 1
	ld.global.f32 	%f137, [%rd22+4];
	add.s64 	%rd24, %rd2, %rd21;
	.loc 1 38 1
	ld.global.f32 	%f138, [%rd24+4];
	add.s64 	%rd26, %rd1, %rd21;
	.loc 1 38 1
	ld.global.f32 	%f139, [%rd26+4];

BB0_4:
	.loc 1 40 1
	setp.gt.s32	%p7, %r3, 0;
	@%p7 bra 	BB0_6;

	mov.f32 	%f142, 0f00000000;
	mov.f32 	%f141, %f142;
	mov.f32 	%f140, %f142;
	bra.uni 	BB0_7;

BB0_6:
	mul.wide.s32 	%rd28, %r17, 4;
	add.s64 	%rd29, %rd3, %rd28;
	.loc 1 42 1
	ld.global.f32 	%f140, [%rd29+-4];
	add.s64 	%rd31, %rd2, %rd28;
	.loc 1 42 1
	ld.global.f32 	%f141, [%rd31+-4];
	add.s64 	%rd33, %rd1, %rd28;
	.loc 1 42 1
	ld.global.f32 	%f142, [%rd33+-4];

BB0_7:
	.loc 1 46 1
	add.f32 	%f19, %f37, %f37;
	.loc 2 3608 3
	div.rn.f32 	%f46, %f35, %f19;
	.loc 1 47 1
	mul.f32 	%f47, %f138, %f138;
	fma.rn.f32 	%f48, %f137, %f137, %f47;
	fma.rn.f32 	%f49, %f139, %f139, %f48;
	.loc 2 3055 10
	sqrt.rn.f32 	%f50, %f49;
	.loc 1 47 59
	setp.eq.f32	%p8, %f50, 0f00000000;
	.loc 1 48 1
	mul.f32 	%f51, %f46, %f38;
	mul.f32 	%f52, %f51, %f6;
	sub.f32 	%f53, %f4, %f52;
	.loc 1 50 1
	mul.f32 	%f54, %f51, %f4;
	add.f32 	%f55, %f6, %f54;
	.loc 1 47 59
	selp.f32	%f56, %f53, %f137, %p8;
	selp.f32	%f57, %f5, %f138, %p8;
	selp.f32	%f58, %f55, %f139, %p8;
	.loc 1 52 1
	mul.f32 	%f59, %f141, %f141;
	fma.rn.f32 	%f60, %f140, %f140, %f59;
	fma.rn.f32 	%f61, %f142, %f142, %f60;
	.loc 2 3055 10
	sqrt.rn.f32 	%f62, %f61;
	.loc 1 52 59
	setp.eq.f32	%p9, %f62, 0f00000000;
	.loc 1 53 1
	add.f32 	%f63, %f4, %f52;
	.loc 1 55 1
	sub.f32 	%f64, %f6, %f54;
	.loc 1 52 59
	selp.f32	%f65, %f63, %f140, %p9;
	selp.f32	%f66, %f5, %f141, %p9;
	selp.f32	%f67, %f64, %f142, %p9;
	.loc 1 58 1
	mul.f32 	%f68, %f38, %f38;
	.loc 2 3608 3
	div.rn.f32 	%f69, %f19, %f68;
	.loc 1 58 31
	sub.f32 	%f70, %f56, %f4;
	sub.f32 	%f71, %f57, %f5;
	sub.f32 	%f72, %f58, %f6;
	sub.f32 	%f73, %f65, %f4;
	sub.f32 	%f74, %f66, %f5;
	sub.f32 	%f75, %f67, %f6;
	add.f32 	%f76, %f70, %f73;
	add.f32 	%f77, %f71, %f74;
	add.f32 	%f78, %f72, %f75;
	.loc 1 58 1
	fma.rn.f32 	%f79, %f69, %f76, %f1;
	fma.rn.f32 	%f20, %f69, %f77, %f2;
	fma.rn.f32 	%f80, %f69, %f78, %f3;
	.loc 1 59 1
	sub.f32 	%f81, %f56, %f65;
	mul.f32 	%f82, %f81, %f35;
	.loc 2 3608 3
	div.rn.f32 	%f83, %f82, %f38;
	.loc 1 59 55
	sub.f32 	%f21, %f80, %f83;
	.loc 1 60 1
	sub.f32 	%f84, %f58, %f67;
	mul.f32 	%f85, %f84, %f35;
	.loc 2 3608 3
	div.rn.f32 	%f86, %f85, %f38;
	.loc 1 60 55
	add.f32 	%f22, %f79, %f86;
	.loc 1 69 1
	add.s32 	%r55, %r2, 1;
	setp.lt.s32	%p10, %r55, %r5;
	@%p10 bra 	BB0_9;

	mov.f32 	%f145, 0f00000000;
	mov.f32 	%f144, %f145;
	mov.f32 	%f143, %f145;
	bra.uni 	BB0_10;

BB0_9:
	.loc 1 70 1
	add.s32 	%r65, %r16, 1;
	mad.lo.s32 	%r70, %r65, %r4, %r3;
	mul.wide.s32 	%rd35, %r70, 4;
	add.s64 	%rd36, %rd3, %rd35;
	.loc 1 71 1
	ld.global.f32 	%f143, [%rd36];
	add.s64 	%rd38, %rd2, %rd35;
	.loc 1 71 1
	ld.global.f32 	%f144, [%rd38];
	add.s64 	%rd40, %rd1, %rd35;
	.loc 1 71 1
	ld.global.f32 	%f145, [%rd40];

BB0_10:
	.loc 1 73 1
	setp.gt.s32	%p11, %r2, 0;
	@%p11 bra 	BB0_12;

	mov.f32 	%f148, 0f00000000;
	mov.f32 	%f147, %f148;
	mov.f32 	%f146, %f148;
	bra.uni 	BB0_13;

BB0_12:
	.loc 1 74 1
	add.s32 	%r84, %r16, -1;
	mad.lo.s32 	%r89, %r84, %r4, %r3;
	mul.wide.s32 	%rd42, %r89, 4;
	add.s64 	%rd43, %rd3, %rd42;
	.loc 1 75 1
	ld.global.f32 	%f146, [%rd43];
	add.s64 	%rd45, %rd2, %rd42;
	.loc 1 75 1
	ld.global.f32 	%f147, [%rd45];
	add.s64 	%rd47, %rd1, %rd42;
	.loc 1 75 1
	ld.global.f32 	%f148, [%rd47];

BB0_13:
	.loc 1 79 1
	mul.f32 	%f93, %f144, %f144;
	fma.rn.f32 	%f94, %f143, %f143, %f93;
	fma.rn.f32 	%f95, %f145, %f145, %f94;
	.loc 2 3055 10
	sqrt.rn.f32 	%f96, %f95;
	.loc 1 79 59
	setp.eq.f32	%p12, %f96, 0f00000000;
	.loc 2 3608 3
	div.rn.f32 	%f97, %f36, %f19;
	.loc 1 81 1
	mul.f32 	%f98, %f97, %f39;
	mul.f32 	%f99, %f98, %f6;
	sub.f32 	%f100, %f5, %f99;
	.loc 1 82 1
	mul.f32 	%f101, %f98, %f5;
	add.f32 	%f102, %f6, %f101;
	.loc 1 79 59
	selp.f32	%f103, %f4, %f143, %p12;
	selp.f32	%f104, %f100, %f144, %p12;
	selp.f32	%f105, %f102, %f145, %p12;
	.loc 1 84 1
	mul.f32 	%f106, %f147, %f147;
	fma.rn.f32 	%f107, %f146, %f146, %f106;
	fma.rn.f32 	%f108, %f148, %f148, %f107;
	.loc 2 3055 10
	sqrt.rn.f32 	%f109, %f108;
	.loc 1 84 59
	setp.eq.f32	%p13, %f109, 0f00000000;
	.loc 1 86 1
	add.f32 	%f110, %f5, %f99;
	.loc 1 87 1
	sub.f32 	%f111, %f6, %f101;
	.loc 1 84 59
	selp.f32	%f112, %f4, %f146, %p13;
	selp.f32	%f113, %f110, %f147, %p13;
	selp.f32	%f114, %f111, %f148, %p13;
	.loc 1 90 1
	mul.f32 	%f115, %f39, %f39;
	.loc 2 3608 3
	div.rn.f32 	%f116, %f19, %f115;
	.loc 1 90 31
	sub.f32 	%f117, %f103, %f4;
	sub.f32 	%f118, %f104, %f5;
	sub.f32 	%f119, %f105, %f6;
	sub.f32 	%f120, %f112, %f4;
	sub.f32 	%f121, %f113, %f5;
	sub.f32 	%f122, %f114, %f6;
	add.f32 	%f123, %f117, %f120;
	add.f32 	%f124, %f118, %f121;
	add.f32 	%f125, %f119, %f122;
	.loc 1 90 1
	fma.rn.f32 	%f126, %f116, %f123, %f22;
	fma.rn.f32 	%f127, %f116, %f124, %f20;
	fma.rn.f32 	%f128, %f116, %f125, %f21;
	.loc 1 91 1
	sub.f32 	%f129, %f104, %f113;
	mul.f32 	%f130, %f129, %f36;
	.loc 2 3608 3
	div.rn.f32 	%f131, %f130, %f39;
	.loc 1 91 55
	sub.f32 	%f132, %f128, %f131;
	.loc 1 92 1
	sub.f32 	%f133, %f105, %f114;
	mul.f32 	%f134, %f133, %f36;
	.loc 2 3608 3
	div.rn.f32 	%f135, %f134, %f39;
	.loc 1 92 55
	add.f32 	%f136, %f127, %f135;
	mul.wide.s32 	%rd49, %r17, 4;
	add.s64 	%rd50, %rd6, %rd49;
	.loc 1 96 1
	st.global.f32 	[%rd50], %f126;
	add.s64 	%rd52, %rd5, %rd49;
	.loc 1 97 1
	st.global.f32 	[%rd52], %f136;
	add.s64 	%rd54, %rd4, %rd49;
	.loc 1 98 1
	st.global.f32 	[%rd54], %f132;

BB0_14:
	.loc 1 99 2
	ret;
}


`
	adddmi_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .f32 adddmi_param_9,
	.param .f32 adddmi_param_10,
	.param .f32 adddmi_param_11,
	.param .f32 adddmi_param_12,
	.param .u32 adddmi_param_13,
	.param .u32 adddmi_param_14,
	.param .u32 adddmi_param_15
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<104>;
	.reg .f32 	%f<149>;
	.reg .s64 	%rd<55>;


	ld.param.u64 	%rd7, [adddmi_param_0];
	ld.param.u64 	%rd8, [adddmi_param_1];
	ld.param.u64 	%rd9, [adddmi_param_2];
	ld.param.u64 	%rd10, [adddmi_param_3];
	ld.param.u64 	%rd11, [adddmi_param_4];
	ld.param.u64 	%rd12, [adddmi_param_5];
	ld.param.f32 	%f35, [adddmi_param_6];
	ld.param.f32 	%f36, [adddmi_param_7];
	ld.param.f32 	%f37, [adddmi_param_9];
	ld.param.f32 	%f38, [adddmi_param_10];
	ld.param.f32 	%f39, [adddmi_param_11];
	ld.param.u32 	%r4, [adddmi_param_13];
	ld.param.u32 	%r5, [adddmi_param_14];
	ld.param.u32 	%r6, [adddmi_param_15];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd9;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd7;
	.loc 1 17 1
	mov.u32 	%r7, %ntid.z;
	mov.u32 	%r8, %ctaid.z;
	mov.u32 	%r9, %tid.z;
	mad.lo.s32 	%r1, %r7, %r8, %r9;
	.loc 1 18 1
	mov.u32 	%r10, %ntid.y;
	mov.u32 	%r11, %ctaid.y;
	mov.u32 	%r12, %tid.y;
	mad.lo.s32 	%r2, %r10, %r11, %r12;
	.loc 1 19 1
	mov.u32 	%r13, %ntid.x;
	mov.u32 	%r14, %ctaid.x;
	mov.u32 	%r15, %tid.x;
	mad.lo.s32 	%r3, %r13, %r14, %r15;
	.loc 1 21 1
	setp.ge.s32	%p1, %r2, %r5;
	setp.ge.s32	%p2, %r3, %r4;
	or.pred  	%p3, %p2, %p1;
	.loc 1 21 1
	setp.ge.s32	%p4, %r1, %r6;
	or.pred  	%p5, %p3, %p4;
	.loc 1 21 1
	@%p5 bra 	BB0_14;

	.loc 1 25 1
	mad.lo.s32 	%r16, %r1, %r5, %r2;
	mad.lo.s32 	%r17, %r16, %r4, %r3;
	mul.wide.s32 	%rd13, %r17, 4;
	add.s64 	%rd14, %rd6, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f1, [%rd14];
	add.s64 	%rd15, %rd5, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f2, [%rd15];
	add.s64 	%rd16, %rd4, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd3, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f4, [%rd17];
	add.s64 	%rd18, %rd2, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f5, [%rd18];
	add.s64 	%rd19, %rd1, %rd13;
	.loc 1 27 1
	ld.global.f32 	%f6, [%rd19];
	.loc 1 36 1
	add.s32 	%r18, %r3, 1;
	setp.lt.s32	%p6, %r18, %r4;
	@%p6 bra 	BB0_3;

	mov.f32 	%f139, 0f00000000;
	mov.f32 	%f138, %f139;
	mov.f32 	%f137, %f139;
	bra.uni 	BB0_4;

BB0_3:
	mul.wide.s32 	%rd21, %r17, 4;
	add.s64 	%rd22, %rd3, %rd21;
	.loc 1 38 1
	ld.global.f32 	%f137, [%rd22+4];
	add.s64 	%rd24, %rd2, %rd21;
	.loc 1 38 1
	ld.global.f32 	%f138, [%rd24+4];
	add.s64 	%rd26, %rd1, %rd21;
	.loc 1 38 1
	ld.global.f32 	%f139, [%rd26+4];

BB0_4:
	.loc 1 40 1
	setp.gt.s32	%p7, %r3, 0;
	@%p7 bra 	BB0_6;

	mov.f32 	%f142, 0f00000000;
	mov.f32 	%f141, %f142;
	mov.f32 	%f140, %f142;
	bra.uni 	BB0_7;

BB0_6:
	mul.wide.s32 	%rd28, %r17, 4;
	add.s64 	%rd29, %rd3, %rd28;
	.loc 1 42 1
	ld.global.f32 	%f140, [%rd29+-4];
	add.s64 	%rd31, %rd2, %rd28;
	.loc 1 42 1
	ld.global.f32 	%f141, [%rd31+-4];
	add.s64 	%rd33, %rd1, %rd28;
	.loc 1 42 1
	ld.global.f32 	%f142, [%rd33+-4];

BB0_7:
	.loc 1 46 1
	add.f32 	%f19, %f37, %f37;
	.loc 2 3608 3
	div.rn.f32 	%f46, %f35, %f19;
	.loc 1 47 1
	mul.f32 	%f47, %f138, %f138;
	fma.rn.f32 	%f48, %f137, %f137, %f47;
	fma.rn.f32 	%f49, %f139, %f139, %f48;
	.loc 2 3055 10
	sqrt.rn.f32 	%f50, %f49;
	.loc 1 47 59
	setp.eq.f32	%p8, %f50, 0f00000000;
	.loc 1 48 1
	mul.f32 	%f51, %f46, %f38;
	mul.f32 	%f52, %f51, %f6;
	sub.f32 	%f53, %f4, %f52;
	.loc 1 50 1
	mul.f32 	%f54, %f51, %f4;
	add.f32 	%f55, %f6, %f54;
	.loc 1 47 59
	selp.f32	%f56, %f53, %f137, %p8;
	selp.f32	%f57, %f5, %f138, %p8;
	selp.f32	%f58, %f55, %f139, %p8;
	.loc 1 52 1
	mul.f32 	%f59, %f141, %f141;
	fma.rn.f32 	%f60, %f140, %f140, %f59;
	fma.rn.f32 	%f61, %f142, %f142, %f60;
	.loc 2 3055 10
	sqrt.rn.f32 	%f62, %f61;
	.loc 1 52 59
	setp.eq.f32	%p9, %f62, 0f00000000;
	.loc 1 53 1
	add.f32 	%f63, %f4, %f52;
	.loc 1 55 1
	sub.f32 	%f64, %f6, %f54;
	.loc 1 52 59
	selp.f32	%f65, %f63, %f140, %p9;
	selp.f32	%f66, %f5, %f141, %p9;
	selp.f32	%f67, %f64, %f142, %p9;
	.loc 1 58 1
	mul.f32 	%f68, %f38, %f38;
	.loc 2 3608 3
	div.rn.f32 	%f69, %f19, %f68;
	.loc 1 58 31
	sub.f32 	%f70, %f56, %f4;
	sub.f32 	%f71, %f57, %f5;
	sub.f32 	%f72, %f58, %f6;
	sub.f32 	%f73, %f65, %f4;
	sub.f32 	%f74, %f66, %f5;
	sub.f32 	%f75, %f67, %f6;
	add.f32 	%f76, %f70, %f73;
	add.f32 	%f77, %f71, %f74;
	add.f32 	%f78, %f72, %f75;
	.loc 1 58 1
	fma.rn.f32 	%f79, %f69, %f76, %f1;
	fma.rn.f32 	%f20, %f69, %f77, %f2;
	fma.rn.f32 	%f80, %f69, %f78, %f3;
	.loc 1 59 1
	sub.f32 	%f81, %f56, %f65;
	mul.f32 	%f82, %f81, %f35;
	.loc 2 3608 3
	div.rn.f32 	%f83, %f82, %f38;
	.loc 1 59 55
	sub.f32 	%f21, %f80, %f83;
	.loc 1 60 1
	sub.f32 	%f84, %f58, %f67;
	mul.f32 	%f85, %f84, %f35;
	.loc 2 3608 3
	div.rn.f32 	%f86, %f85, %f38;
	.loc 1 60 55
	add.f32 	%f22, %f79, %f86;
	.loc 1 69 1
	add.s32 	%r55, %r2, 1;
	setp.lt.s32	%p10, %r55, %r5;
	@%p10 bra 	BB0_9;

	mov.f32 	%f145, 0f00000000;
	mov.f32 	%f144, %f145;
	mov.f32 	%f143, %f145;
	bra.uni 	BB0_10;

BB0_9:
	.loc 1 70 1
	add.s32 	%r65, %r16, 1;
	mad.lo.s32 	%r70, %r65, %r4, %r3;
	mul.wide.s32 	%rd35, %r70, 4;
	add.s64 	%rd36, %rd3, %rd35;
	.loc 1 71 1
	ld.global.f32 	%f143, [%rd36];
	add.s64 	%rd38, %rd2, %rd35;
	.loc 1 71 1
	ld.global.f32 	%f144, [%rd38];
	add.s64 	%rd40, %rd1, %rd35;
	.loc 1 71 1
	ld.global.f32 	%f145, [%rd40];

BB0_10:
	.loc 1 73 1
	setp.gt.s32	%p11, %r2, 0;
	@%p11 bra 	BB0_12;

	mov.f32 	%f148, 0f00000000;
	mov.f32 	%f147, %f148;
	mov.f32 	%f146, %f148;
	bra.uni 	BB0_13;

BB0_12:
	.loc 1 74 1
	add.s32 	%r84, %r16, -1;
	mad.lo.s32 	%r89, %r84, %r4, %r3;
	mul.wide.s32 	%rd42, %r89, 4;
	add.s64 	%rd43, %rd3, %rd42;
	.loc 1 75 1
	ld.global.f32 	%f146, [%rd43];
	add.s64 	%rd45, %rd2, %rd42;
	.loc 1 75 1
	ld.global.f32 	%f147, [%rd45];
	add.s64 	%rd47, %rd1, %rd42;
	.loc 1 75 1
	ld.global.f32 	%f148, [%rd47];

BB0_13:
	.loc 1 79 1
	mul.f32 	%f93, %f144, %f144;
	fma.rn.f32 	%f94, %f143, %f143, %f93;
	fma.rn.f32 	%f95, %f145, %f145, %f94;
	.loc 2 3055 10
	sqrt.rn.f32 	%f96, %f95;
	.loc 1 79 59
	setp.eq.f32	%p12, %f96, 0f00000000;
	.loc 2 3608 3
	div.rn.f32 	%f97, %f36, %f19;
	.loc 1 81 1
	mul.f32 	%f98, %f97, %f39;
	mul.f32 	%f99, %f98, %f6;
	sub.f32 	%f100, %f5, %f99;
	.loc 1 82 1
	mul.f32 	%f101, %f98, %f5;
	add.f32 	%f102, %f6, %f101;
	.loc 1 79 59
	selp.f32	%f103, %f4, %f143, %p12;
	selp.f32	%f104, %f100, %f144, %p12;
	selp.f32	%f105, %f102, %f145, %p12;
	.loc 1 84 1
	mul.f32 	%f106, %f147, %f147;
	fma.rn.f32 	%f107, %f146, %f146, %f106;
	fma.rn.f32 	%f108, %f148, %f148, %f107;
	.loc 2 3055 10
	sqrt.rn.f32 	%f109, %f108;
	.loc 1 84 59
	setp.eq.f32	%p13, %f109, 0f00000000;
	.loc 1 86 1
	add.f32 	%f110, %f5, %f99;
	.loc 1 87 1
	sub.f32 	%f111, %f6, %f101;
	.loc 1 84 59
	selp.f32	%f112, %f4, %f146, %p13;
	selp.f32	%f113, %f110, %f147, %p13;
	selp.f32	%f114, %f111, %f148, %p13;
	.loc 1 90 1
	mul.f32 	%f115, %f39, %f39;
	.loc 2 3608 3
	div.rn.f32 	%f116, %f19, %f115;
	.loc 1 90 31
	sub.f32 	%f117, %f103, %f4;
	sub.f32 	%f118, %f104, %f5;
	sub.f32 	%f119, %f105, %f6;
	sub.f32 	%f120, %f112, %f4;
	sub.f32 	%f121, %f113, %f5;
	sub.f32 	%f122, %f114, %f6;
	add.f32 	%f123, %f117, %f120;
	add.f32 	%f124, %f118, %f121;
	add.f32 	%f125, %f119, %f122;
	.loc 1 90 1
	fma.rn.f32 	%f126, %f116, %f123, %f22;
	fma.rn.f32 	%f127, %f116, %f124, %f20;
	fma.rn.f32 	%f128, %f116, %f125, %f21;
	.loc 1 91 1
	sub.f32 	%f129, %f104, %f113;
	mul.f32 	%f130, %f129, %f36;
	.loc 2 3608 3
	div.rn.f32 	%f131, %f130, %f39;
	.loc 1 91 55
	sub.f32 	%f132, %f128, %f131;
	.loc 1 92 1
	sub.f32 	%f133, %f105, %f114;
	mul.f32 	%f134, %f133, %f36;
	.loc 2 3608 3
	div.rn.f32 	%f135, %f134, %f39;
	.loc 1 92 55
	add.f32 	%f136, %f127, %f135;
	mul.wide.s32 	%rd49, %r17, 4;
	add.s64 	%rd50, %rd6, %rd49;
	.loc 1 96 1
	st.global.f32 	[%rd50], %f126;
	add.s64 	%rd52, %rd5, %rd49;
	.loc 1 97 1
	st.global.f32 	[%rd52], %f136;
	add.s64 	%rd54, %rd4, %rd49;
	.loc 1 98 1
	st.global.f32 	[%rd54], %f132;

BB0_14:
	.loc 1 99 2
	ret;
}


`
	adddmi_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .f32 adddmi_param_9,
	.param .f32 adddmi_param_10,
	.param .f32 adddmi_param_11,
	.param .f32 adddmi_param_12,
	.param .u32 adddmi_param_13,
	.param .u32 adddmi_param_14,
	.param .u32 adddmi_param_15
)
{
	.reg .pred 	%p<14>;
	.reg .s32 	%r<104>;
	.reg .f32 	%f<149>;
	.reg .s64 	%rd<55>;


	ld.param.u64 	%rd7, [adddmi_param_0];
	ld.param.u64 	%rd8, [adddmi_param_1];
	ld.param.u64 	%rd9, [adddmi_param_2];
	ld.param.u64 	%rd10, [adddmi_param_3];
	ld.param.u64 	%rd11, [adddmi_param_4];
	ld.param.u64 	%rd12, [adddmi_param_5];
	ld.param.f32 	%f35, [adddmi_param_6];
	ld.param.f32 	%f36, [adddmi_param_7];
	ld.param.f32 	%f37, [adddmi_param_9];
	ld.param.f32 	%f38, [adddmi_param_10];
	ld.param.f32 	%f39, [adddmi_param_11];
	ld.param.u32 	%r4, [adddmi_param_13];
	ld.param.u32 	%r5, [adddmi_param_14];
	ld.param.u32 	%r6, [adddmi_param_15];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd9;
	cvta.to.global.u64 	%rd5, %rd8;
	cvta.to.global.u64 	%rd6, %rd7;
	.loc 1 17 1
	mov.u32 	%r7, %ntid.z;
	mov.u32 	%r8, %ctaid.z;
	mov.u32 	%r9, %tid.z;
	mad.lo.s32 	%r1, %r7, %r8, %r9;
	.loc 1 18 1
	mov.u32 	%r10, %ntid.y;
	mov.u32 	%r11, %ctaid.y;
	mov.u32 	%r12, %tid.y;
	mad.lo.s32 	%r2, %r10, %r11, %r12;
	.loc 1 19 1
	mov.u32 	%r13, %ntid.x;
	mov.u32 	%r14, %ctaid.x;
	mov.u32 	%r15, %tid.x;
	mad.lo.s32 	%r3, %r13, %r14, %r15;
	.loc 1 21 1
	setp.ge.s32	%p1, %r2, %r5;
	setp.ge.s32	%p2, %r3, %r4;
	or.pred  	%p3, %p2, %p1;
	.loc 1 21 1
	setp.ge.s32	%p4, %r1, %r6;
	or.pred  	%p5, %p3, %p4;
	.loc 1 21 1
	@%p5 bra 	BB2_14;

	.loc 1 25 1
	mad.lo.s32 	%r16, %r1, %r5, %r2;
	mad.lo.s32 	%r17, %r16, %r4, %r3;
	mul.wide.s32 	%rd13, %r17, 4;
	add.s64 	%rd14, %rd6, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f1, [%rd14];
	add.s64 	%rd15, %rd5, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f2, [%rd15];
	add.s64 	%rd16, %rd4, %rd13;
	.loc 1 26 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd3, %rd13;
	.loc 1 27 1
	ld.global.nc.f32 	%f4, [%rd17];
	add.s64 	%rd18, %rd2, %rd13;
	.loc 1 27 1
	ld.global.nc.f32 	%f5, [%rd18];
	add.s64 	%rd19, %rd1, %rd13;
	.loc 1 27 1
	ld.global.nc.f32 	%f6, [%rd19];
	.loc 1 36 1
	add.s32 	%r18, %r3, 1;
	setp.lt.s32	%p6, %r18, %r4;
	@%p6 bra 	BB2_3;

	mov.f32 	%f139, 0f00000000;
	mov.f32 	%f138, %f139;
	mov.f32 	%f137, %f139;
	bra.uni 	BB2_4;

BB2_3:
	mul.wide.s32 	%rd21, %r17, 4;
	add.s64 	%rd22, %rd3, %rd21;
	.loc 1 38 1
	ld.global.nc.f32 	%f137, [%rd22+4];
	add.s64 	%rd24, %rd2, %rd21;
	.loc 1 38 1
	ld.global.nc.f32 	%f138, [%rd24+4];
	add.s64 	%rd26, %rd1, %rd21;
	.loc 1 38 1
	ld.global.nc.f32 	%f139, [%rd26+4];

BB2_4:
	.loc 1 40 1
	setp.gt.s32	%p7, %r3, 0;
	@%p7 bra 	BB2_6;

	mov.f32 	%f142, 0f00000000;
	mov.f32 	%f141, %f142;
	mov.f32 	%f140, %f142;
	bra.uni 	BB2_7;

BB2_6:
	mul.wide.s32 	%rd28, %r17, 4;
	add.s64 	%rd29, %rd3, %rd28;
	.loc 1 42 1
	ld.global.nc.f32 	%f140, [%rd29+-4];
	add.s64 	%rd31, %rd2, %rd28;
	.loc 1 42 1
	ld.global.nc.f32 	%f141, [%rd31+-4];
	add.s64 	%rd33, %rd1, %rd28;
	.loc 1 42 1
	ld.global.nc.f32 	%f142, [%rd33+-4];

BB2_7:
	.loc 1 46 1
	add.f32 	%f19, %f37, %f37;
	.loc 3 3608 3
	div.rn.f32 	%f46, %f35, %f19;
	.loc 1 47 1
	mul.f32 	%f47, %f138, %f138;
	fma.rn.f32 	%f48, %f137, %f137, %f47;
	fma.rn.f32 	%f49, %f139, %f139, %f48;
	.loc 3 3055 10
	sqrt.rn.f32 	%f50, %f49;
	.loc 1 47 59
	setp.eq.f32	%p8, %f50, 0f00000000;
	.loc 1 48 1
	mul.f32 	%f51, %f46, %f38;
	mul.f32 	%f52, %f51, %f6;
	sub.f32 	%f53, %f4, %f52;
	.loc 1 50 1
	mul.f32 	%f54, %f51, %f4;
	add.f32 	%f55, %f6, %f54;
	.loc 1 47 59
	selp.f32	%f56, %f53, %f137, %p8;
	selp.f32	%f57, %f5, %f138, %p8;
	selp.f32	%f58, %f55, %f139, %p8;
	.loc 1 52 1
	mul.f32 	%f59, %f141, %f141;
	fma.rn.f32 	%f60, %f140, %f140, %f59;
	fma.rn.f32 	%f61, %f142, %f142, %f60;
	.loc 3 3055 10
	sqrt.rn.f32 	%f62, %f61;
	.loc 1 52 59
	setp.eq.f32	%p9, %f62, 0f00000000;
	.loc 1 53 1
	add.f32 	%f63, %f4, %f52;
	.loc 1 55 1
	sub.f32 	%f64, %f6, %f54;
	.loc 1 52 59
	selp.f32	%f65, %f63, %f140, %p9;
	selp.f32	%f66, %f5, %f141, %p9;
	selp.f32	%f67, %f64, %f142, %p9;
	.loc 1 58 1
	mul.f32 	%f68, %f38, %f38;
	.loc 3 3608 3
	div.rn.f32 	%f69, %f19, %f68;
	.loc 1 58 31
	sub.f32 	%f70, %f56, %f4;
	sub.f32 	%f71, %f57, %f5;
	sub.f32 	%f72, %f58, %f6;
	sub.f32 	%f73, %f65, %f4;
	sub.f32 	%f74, %f66, %f5;
	sub.f32 	%f75, %f67, %f6;
	add.f32 	%f76, %f70, %f73;
	add.f32 	%f77, %f71, %f74;
	add.f32 	%f78, %f72, %f75;
	.loc 1 58 1
	fma.rn.f32 	%f79, %f69, %f76, %f1;
	fma.rn.f32 	%f20, %f69, %f77, %f2;
	fma.rn.f32 	%f80, %f69, %f78, %f3;
	.loc 1 59 1
	sub.f32 	%f81, %f56, %f65;
	mul.f32 	%f82, %f81, %f35;
	.loc 3 3608 3
	div.rn.f32 	%f83, %f82, %f38;
	.loc 1 59 55
	sub.f32 	%f21, %f80, %f83;
	.loc 1 60 1
	sub.f32 	%f84, %f58, %f67;
	mul.f32 	%f85, %f84, %f35;
	.loc 3 3608 3
	div.rn.f32 	%f86, %f85, %f38;
	.loc 1 60 55
	add.f32 	%f22, %f79, %f86;
	.loc 1 69 1
	add.s32 	%r55, %r2, 1;
	setp.lt.s32	%p10, %r55, %r5;
	@%p10 bra 	BB2_9;

	mov.f32 	%f145, 0f00000000;
	mov.f32 	%f144, %f145;
	mov.f32 	%f143, %f145;
	bra.uni 	BB2_10;

BB2_9:
	.loc 1 70 1
	add.s32 	%r65, %r16, 1;
	mad.lo.s32 	%r70, %r65, %r4, %r3;
	mul.wide.s32 	%rd35, %r70, 4;
	add.s64 	%rd36, %rd3, %rd35;
	.loc 1 71 1
	ld.global.nc.f32 	%f143, [%rd36];
	add.s64 	%rd38, %rd2, %rd35;
	.loc 1 71 1
	ld.global.nc.f32 	%f144, [%rd38];
	add.s64 	%rd40, %rd1, %rd35;
	.loc 1 71 1
	ld.global.nc.f32 	%f145, [%rd40];

BB2_10:
	.loc 1 73 1
	setp.gt.s32	%p11, %r2, 0;
	@%p11 bra 	BB2_12;

	mov.f32 	%f148, 0f00000000;
	mov.f32 	%f147, %f148;
	mov.f32 	%f146, %f148;
	bra.uni 	BB2_13;

BB2_12:
	.loc 1 74 1
	add.s32 	%r84, %r16, -1;
	mad.lo.s32 	%r89, %r84, %r4, %r3;
	mul.wide.s32 	%rd42, %r89, 4;
	add.s64 	%rd43, %rd3, %rd42;
	.loc 1 75 1
	ld.global.nc.f32 	%f146, [%rd43];
	add.s64 	%rd45, %rd2, %rd42;
	.loc 1 75 1
	ld.global.nc.f32 	%f147, [%rd45];
	add.s64 	%rd47, %rd1, %rd42;
	.loc 1 75 1
	ld.global.nc.f32 	%f148, [%rd47];

BB2_13:
	.loc 1 79 1
	mul.f32 	%f93, %f144, %f144;
	fma.rn.f32 	%f94, %f143, %f143, %f93;
	fma.rn.f32 	%f95, %f145, %f145, %f94;
	.loc 3 3055 10
	sqrt.rn.f32 	%f96, %f95;
	.loc 1 79 59
	setp.eq.f32	%p12, %f96, 0f00000000;
	.loc 3 3608 3
	div.rn.f32 	%f97, %f36, %f19;
	.loc 1 81 1
	mul.f32 	%f98, %f97, %f39;
	mul.f32 	%f99, %f98, %f6;
	sub.f32 	%f100, %f5, %f99;
	.loc 1 82 1
	mul.f32 	%f101, %f98, %f5;
	add.f32 	%f102, %f6, %f101;
	.loc 1 79 59
	selp.f32	%f103, %f4, %f143, %p12;
	selp.f32	%f104, %f100, %f144, %p12;
	selp.f32	%f105, %f102, %f145, %p12;
	.loc 1 84 1
	mul.f32 	%f106, %f147, %f147;
	fma.rn.f32 	%f107, %f146, %f146, %f106;
	fma.rn.f32 	%f108, %f148, %f148, %f107;
	.loc 3 3055 10
	sqrt.rn.f32 	%f109, %f108;
	.loc 1 84 59
	setp.eq.f32	%p13, %f109, 0f00000000;
	.loc 1 86 1
	add.f32 	%f110, %f5, %f99;
	.loc 1 87 1
	sub.f32 	%f111, %f6, %f101;
	.loc 1 84 59
	selp.f32	%f112, %f4, %f146, %p13;
	selp.f32	%f113, %f110, %f147, %p13;
	selp.f32	%f114, %f111, %f148, %p13;
	.loc 1 90 1
	mul.f32 	%f115, %f39, %f39;
	.loc 3 3608 3
	div.rn.f32 	%f116, %f19, %f115;
	.loc 1 90 31
	sub.f32 	%f117, %f103, %f4;
	sub.f32 	%f118, %f104, %f5;
	sub.f32 	%f119, %f105, %f6;
	sub.f32 	%f120, %f112, %f4;
	sub.f32 	%f121, %f113, %f5;
	sub.f32 	%f122, %f114, %f6;
	add.f32 	%f123, %f117, %f120;
	add.f32 	%f124, %f118, %f121;
	add.f32 	%f125, %f119, %f122;
	.loc 1 90 1
	fma.rn.f32 	%f126, %f116, %f123, %f22;
	fma.rn.f32 	%f127, %f116, %f124, %f20;
	fma.rn.f32 	%f128, %f116, %f125, %f21;
	.loc 1 91 1
	sub.f32 	%f129, %f104, %f113;
	mul.f32 	%f130, %f129, %f36;
	.loc 3 3608 3
	div.rn.f32 	%f131, %f130, %f39;
	.loc 1 91 55
	sub.f32 	%f132, %f128, %f131;
	.loc 1 92 1
	sub.f32 	%f133, %f105, %f114;
	mul.f32 	%f134, %f133, %f36;
	.loc 3 3608 3
	div.rn.f32 	%f135, %f134, %f39;
	.loc 1 92 55
	add.f32 	%f136, %f127, %f135;
	mul.wide.s32 	%rd49, %r17, 4;
	add.s64 	%rd50, %rd6, %rd49;
	.loc 1 96 1
	st.global.f32 	[%rd50], %f126;
	add.s64 	%rd52, %rd5, %rd49;
	.loc 1 97 1
	st.global.f32 	[%rd52], %f136;
	add.s64 	%rd54, %rd4, %rd49;
	.loc 1 98 1
	st.global.f32 	[%rd54], %f132;

BB2_14:
	.loc 1 99 2
	ret;
}


`
)
